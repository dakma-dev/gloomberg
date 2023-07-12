package gloomberg

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
	"github.com/spf13/viper"
)

// ErrWalletBalance given for issues while fetching wallet balances.
var ErrWalletBalance = fmt.Errorf("error fetching wallet balance")

var (
	listStyle = lipgloss.NewStyle()
	// Border(lipgloss.NormalBorder(), false, true, false, false).
	// BorderForeground(style.Subtle)
	// MarginRight(0).

	itemStyle = lipgloss.NewStyle().Padding(0, 2)
	listItem  = itemStyle.Render
)

type Stats struct {
	gb           *Gloomberg
	wallets      *wallet.Wallets
	providerPool *provider.Pool
	rdb          rueidis.Client

	interval  time.Duration
	timeframe time.Duration

	RecentEvents    mapset.Set[*degendb.RecentEvent]
	RecentOwnEvents mapset.Set[*degendb.ParsedEvent]

	gasTicker *time.Ticker

	NewLogs        uint64
	NewListings    uint64
	EventsToFormat uint64
	OutputLines    uint64
}

func NewStats(gb *Gloomberg, gasTicker *time.Ticker, wallets *wallet.Wallets, providerPool *provider.Pool, rdb rueidis.Client) *Stats {
	stats := &Stats{
		gb:           gb,
		wallets:      wallets,
		providerPool: providerPool,
		rdb:          rdb,

		RecentEvents:    mapset.NewSet[*degendb.RecentEvent](),
		RecentOwnEvents: mapset.NewSet[*degendb.ParsedEvent](),

		gasTicker: gasTicker,

		interval:  viper.GetDuration("ticker.statsbox"),
		timeframe: viper.GetDuration("stats.timeframe"),
	}

	return stats
}

func (s *Stats) AddEvent(eventType degendb.EventType, amountTokens int64, value *big.Int) {
	s.RecentEvents.Add(&degendb.RecentEvent{
		Timestamp:    time.Now(),
		Type:         eventType,
		AmountTokens: uint64(amountTokens),
		AmountWei:    value,
	})
}

func (s *Stats) eventslastTimeframe(eventType degendb.EventType) uint64 {
	count := uint64(0)

	for _, event := range s.RecentEvents.ToSlice() {
		if event.Type == eventType && time.Since(event.Timestamp) < s.timeframe {
			count++
		}
	}

	return count
}

func (s *Stats) volumeLastTimeframe() *big.Int {
	volume := big.NewInt(0)

	for _, event := range s.RecentEvents.ToSlice() {
		if event.Type == degendb.Sale && time.Since(event.Timestamp) < s.timeframe {
			volume = volume.Add(volume, event.AmountWei)
		}
	}

	return volume
}

func (s *Stats) UpdateBalances() (*wallet.Wallets, error) {
	gbl.Log.Debugf("updating wallet balances...")

	balances, err := external.GetBalances(s.wallets)
	if err != nil || balances == nil {
		gbl.Log.Warn("❌ error while fetching wallet balances")

		return nil, ErrWalletBalance
	}

	if viper.GetBool("log.debug") {
		for _, balance := range balances {
			gbl.Log.Debugf("UpdateBalances| %+v\n", balance)
		}
	}

	for _, balance := range balances {
		walletAddress := common.HexToAddress(balance.Account)

		// init with ETH Balance
		balanceTotalWei := balance.BalanceETH

		// add WETH Balance
		balanceTotalWei = big.NewInt(0).Add(balanceTotalWei, balance.BalanceWETH)

		// add BlurPool Balance
		if balance.BalanceBlurPool != nil {
			balanceTotalWei = big.NewInt(0).Add(balanceTotalWei, balance.BalanceBlurPool)
		}

		gbl.Log.Debugf("%s: %6.3fΞ total || %6.3f ETH | %6.3f WETH | %6.3f BlurPool", balance.Account, balanceTotalWei, utils.WeiToEther(balance.BalanceETH), utils.WeiToEther(balance.BalanceWETH), utils.WeiToEther(balance.BalanceBlurPool))

		(*s.wallets)[walletAddress].BalanceBefore = (*s.wallets)[walletAddress].Balance
		(*s.wallets)[walletAddress].Balance = balanceTotalWei

		trendIndicator := style.CreateTrendIndicator(
			float64((*s.wallets)[walletAddress].BalanceBefore.Int64()),
			float64((*s.wallets)[walletAddress].Balance.Int64()),
		)

		(*s.wallets)[walletAddress].BalanceTrend = trendIndicator.String()

		gbl.Log.Debugf("  %s balance: %s %6.3f", balance.Account, trendIndicator, utils.WeiToEther((*s.wallets)[walletAddress].Balance))
	}

	return s.wallets, nil
}

func (s *Stats) Print(queueOutput chan string) {
	var (
		formattedStatsLists string

		statsLists []string
	)

	if viper.GetBool("stats.balances") {
		_, err := s.UpdateBalances()
		if err != nil {
			gbl.Log.Warn("❌ error while updating w balances")
		}
	}

	// new
	statsLists = []string{}
	statsLists = append(statsLists, s.getPrimaryStatsLists()...)

	maxWalletNameLength := 0.0
	for _, w := range *s.wallets {
		maxWalletNameLength = math.Max(maxWalletNameLength, float64(len(w.Name)))
	}

	if walletBalancesList := s.getWalletStatsList(int(maxWalletNameLength)); len(walletBalancesList) > 0 {
		statsLists = append(statsLists, listStyle.Render(lipgloss.JoinVertical(lipgloss.Left, walletBalancesList...)))
	}

	if s.RecentOwnEvents.Cardinality() > 0 {
		eventsList := listStyle // .Copy().UnsetWidth().PaddingLeft(0).Render
		statsLists = append(statsLists, eventsList.Render(lipgloss.JoinVertical(lipgloss.Left, s.getOwnEventsHistoryList()...)))
	}

	formattedStatsLists = lipgloss.JoinHorizontal(lipgloss.Top, statsLists...)

	if s.gasTicker != nil {
		s.gasTicker.Reset(viper.GetDuration("ticker.gasline"))
	}

	queueOutput <- "\n" + formattedStatsLists + "\n"
}

func (s *Stats) getPrimaryStatsLists() []string {
	// first column
	var firstColumn []string

	// gas
	if gasInfo, err := s.providerPool.GetCurrentGasInfo(); err == nil && gasInfo != nil {
		// gas info
		if gasInfo.GasPriceWei.Cmp(big.NewInt(0)) > 0 {
			gasPriceGwei, _ := utils.WeiToGwei(gasInfo.GasPriceWei).Float64()
			gasPrice := uint64(math.Ceil(gasPriceGwei))
			// gasTip, _ := nodes.WeiToGwei(gasInfo.GasTipWei).Uint64()

			atomic.StoreUint64(&s.gb.CurrentGasPriceGwei, gasPrice)

			label := style.DarkGrayStyle.Render("   gas")
			value := style.LightGrayStyle.Render(fmt.Sprintf("%3d", gasPrice))

			firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", label, value)), listItem("")}...)
		}
	} else if viper.IsSet("api_keys.etherscan") {
		label := style.DarkGrayStyle.Render("  gas")
		value := style.LightGrayStyle.Render(fmt.Sprintf("%3d", external.GetEstimatedGasPrice()))

		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", label, value)), listItem("")}...)
	}

	//
	// (per minute) stats

	// get volume from recent events
	volumeWei := s.volumeLastTimeframe()

	// time used for normalizing to per minute (stats.timeframe > runtime ? stats.timeframe : runtime)
	timeframeMinutes := s.getTimeframeMinutes()

	// show a wait indicator if we are not running for the full timeframe yet
	unitOrWait := "m"
	if !s.runningForFullTimeframe() {
		unitOrWait = "⏳"
	}

	// until runtime > timeframe, our data is not yet accurate -> darken the values
	valueStyle := style.GrayStyle

	// assert volume and timeframeMinutes are > 0
	if volume, _ := utils.WeiToEther(volumeWei).Float64(); volume*timeframeMinutes > 0 {
		volumePerMin := volume / timeframeMinutes

		volumeLabel := style.DarkGrayStyle.Render("Ξ /" + unitOrWait)
		volumeStyle := valueStyle.Copy()

		switch {
		case !s.runningForFullTimeframe():
			volumeStyle = style.DarkGrayStyle
		case volumePerMin > 100.0:
			volumeStyle = style.PinkBoldStyle
		case volumePerMin > 42:
			volumeStyle = style.AlmostWhiteStyle
		case volumePerMin > 13:
			volumeStyle = style.LightGrayStyle
		}

		fmtVolumePerMin := volumeStyle.Render(fmt.Sprintf("%6.2f", volumePerMin))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s%s", fmtVolumePerMin, volumeLabel))}...)
	}

	if salesCount := s.eventslastTimeframe(degendb.Sale); salesCount > 0 {
		salesPerMin := int(float64(salesCount) / timeframeMinutes)

		salesLabel := style.DarkGrayStyle.Render("s/" + unitOrWait)
		salesStyle := valueStyle.Copy()

		switch {
		case !s.runningForFullTimeframe():
			salesStyle = style.DarkGrayStyle
		case salesPerMin > 300:
			salesStyle = style.PinkBoldStyle
		case salesPerMin > 200:
			salesStyle = style.AlmostWhiteStyle
		case salesPerMin > 100:
			salesStyle = style.LightGrayStyle
		}

		fmtSalesPerMin := salesStyle.Render(fmt.Sprintf("%6d", salesPerMin))

		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", fmtSalesPerMin, salesLabel))}...)
	}

	if mintsCount := s.eventslastTimeframe(degendb.Mint); mintsCount > 0 {
		mintsLabel := style.DarkGrayStyle.Render("m/" + unitOrWait)

		mintsPerMin := int(float64(mintsCount) / timeframeMinutes)

		mintsStyle := valueStyle.Copy()

		switch {
		case !s.runningForFullTimeframe():
			mintsStyle = style.DarkGrayStyle
		case mintsPerMin > viper.GetInt("stats.high_volume.mints.threshold"):
			go s.highVolumeMint()

			fallthrough
		case mintsPerMin > 500:
			mintsStyle = style.PinkBoldStyle
		case mintsPerMin > 300:
			mintsStyle = style.AlmostWhiteStyle
		case mintsPerMin > 150:
			mintsStyle = style.LightGrayStyle
		}

		fmtMintsPerMin := mintsStyle.Render(fmt.Sprintf("%6d", mintsPerMin))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", fmtMintsPerMin, mintsLabel))}...)
	}

	//
	// second column
	var secondcolumn []string

	// min price
	if minPrice := viper.GetFloat64("show.min_value"); minPrice > 0.0 {
		label := style.DarkGrayStyle.Render("min price")
		value := style.GrayStyle.Render(fmt.Sprint(fmt.Sprintf("%6.2f", minPrice), style.DarkGrayStyle.Render("Ξ")))

		secondcolumn = append(secondcolumn, []string{listItem(fmt.Sprintf("%s %s", label, value)), listItem("")}...)
	}

	// redis stats
	if viper.GetBool("redis.enabled") {
		if s.rdb != nil {
			dbSize, err := s.rdb.Do(context.TODO(), s.rdb.B().Dbsize().Build()).AsInt64()
			if err != nil {
				gbl.Log.Warnf("failed to get redis dbsize: %v", err)
			}

			dbInfo, err := s.rdb.Do(context.TODO(), s.rdb.B().Info().Section("stats").Build()).ToString()
			if err != nil {
				gbl.Log.Warnf("failed to get redis dbsize: %v", err)
			}

			// cache hitrate
			var keyspaceHits, keyspaceMisses int64

			for _, stat := range strings.Split(dbInfo, "\r\n") {
				if rawKeyspaceHits := strings.TrimPrefix(stat, "keyspace_hits:"); rawKeyspaceHits != stat {
					keyspaceHits, err = strconv.ParseInt(rawKeyspaceHits, 10, 64)
					if err != nil {
						gbl.Log.Warnf("failed to parse keyspace_hits: %v", err)
					}

					log.Debugf("keyspaceHits: %+v", keyspaceHits)
				}

				if rawKeyspaceMisses := strings.TrimPrefix(stat, "keyspace_misses:"); rawKeyspaceMisses != stat {
					keyspaceMisses, err = strconv.ParseInt(rawKeyspaceMisses, 10, 64)
					if err != nil {
						gbl.Log.Warnf("failed to parse keyspaceMisses: %v", err)
					}

					// we're done
					break
				}
			}

			// calculate hitrate
			hitrate := (float64(keyspaceHits) / float64(keyspaceHits+keyspaceMisses)) * 100
			log.Debugf("keyspace_hits: %+v | keyspace_misses: %+v | hitrate: %+v", keyspaceHits, keyspaceMisses, hitrate)

			// cache size
			namesLabel := style.DarkGrayStyle.Render("  cache")
			namesValue := style.GrayStyle.Render(fmt.Sprintf("%9d", dbSize))

			// hitrate
			hitrateLabel := style.DarkGrayStyle.Render("hitrate")
			hitrateValue := fmt.Sprint(style.GrayStyle.Render(fmt.Sprintf("%8.2f", hitrate)), style.DarkGrayStyle.Render("%"))

			// add to second column
			secondcolumn = append(secondcolumn, []string{
				listItem(fmt.Sprintf("%s %s", namesLabel, namesValue)),
				// listItem(fmt.Sprintf("%s %s", salesLabel, salesValue)),
				listItem(fmt.Sprintf("%s %s", hitrateLabel, hitrateValue)),
			}...)
		}
	}

	// running for
	labelRunningFor := style.DarkGrayStyle.Render("running")
	valueRunningFor := style.GrayStyle.Copy().Width(9).Align(lipgloss.Right).Render(time.Since(internal.RunningSince).Truncate(time.Second).String())

	secondcolumn = append(secondcolumn, []string{listItem(labelRunningFor + " " + valueRunningFor)}...)

	// running since
	labelRunningSince := style.DarkGrayStyle.Render("  since")
	valueRunningSince := style.GrayStyle.Copy().Width(9).Align(lipgloss.Right).Render(internal.RunningSince.Format("15:04"))

	secondcolumn = append(secondcolumn, []string{listItem(fmt.Sprintf("%s %s", labelRunningSince, valueRunningSince))}...)

	// combine lists
	statsOutput := []string{listStyle.Copy().Render(lipgloss.JoinVertical(lipgloss.Left, firstColumn...))}

	if len(secondcolumn) > 0 {
		statsOutput = append(statsOutput, listStyle.Copy().Render(lipgloss.JoinVertical(lipgloss.Left, secondcolumn...)))
	}

	return statsOutput
}

func (s *Stats) getWalletStatsList(maxWalletNameLength int) []string {
	wallets := s.wallets.SortByBalance()

	numberOfWalletsToShow := int(math.Min(float64(viper.GetInt("stats.lines")), float64(len(wallets))))

	walletsList := make([]string, 0)

	for _, w := range wallets[:numberOfWalletsToShow] {
		balanceEther, _ := utils.WeiToEther(w.Balance).Float64()
		balanceRounded := math.Floor(balanceEther*100.0) / 100.0
		balance := fmt.Sprint(style.LightGrayStyle.Render(fmt.Sprintf("%5.2f", balanceRounded)), style.GrayStyle.Render("Ξ"))
		walletBalance := fmt.Sprintf("%s %s %s", w.ColoredName(maxWalletNameLength), style.DarkGrayStyle.Render(w.BalanceTrend), balance)
		walletsList = append(walletsList, listItem(walletBalance))
	}

	return walletsList
}

func (s *Stats) getOwnEventsHistoryList() []string {
	eventsList := make([]string, 0)

	if s.RecentOwnEvents.Cardinality() == 0 {
		gbl.Log.Debugf("no events to show")

		return eventsList
	}

	// cleanup (maybe replace this by not inserting events that are not shown anyways)
	historyEvents := make([]*degendb.ParsedEvent, 0)

	// for idx, event := range s.EventHistory {
	for idx, event := range s.RecentOwnEvents.ToSlice() {
		if event == nil {
			gbl.Log.Debugf("␀ event is nil: %d\n", idx)

			continue
		}

		historyEvents = append(historyEvents, event)
	}

	sort.Slice(historyEvents, func(i, j int) bool { return historyEvents[i].ReceivedAt.Before(historyEvents[j].ReceivedAt) })

	numberOfOwnEvents := len(historyEvents)
	numberOfShownEvents := int(math.Min(float64(viper.GetInt("stats.lines")), float64(numberOfOwnEvents)))
	firstEventShown := numberOfOwnEvents - numberOfShownEvents

	for _, event := range historyEvents[firstEventShown:] {
		if len(eventsList) >= numberOfShownEvents {
			break
		}

		if len(event.TransferredCollections) == 0 {
			continue
		}

		collectionStyle := lipgloss.NewStyle().Foreground(event.TransferredCollections[0].Colors.Primary)

		timeAgo := time.Since(event.ReceivedAt)
		statsboxEpoch := viper.GetDuration("ticker.statsbox")

		rowStyle := style.DarkGrayStyle
		printFaint := false

		switch {
		case timeAgo < statsboxEpoch:
			rowStyle = style.BoldStyle
		case timeAgo < 2*statsboxEpoch:
			rowStyle = style.DarkWhiteStyle
		case timeAgo < 4*statsboxEpoch:
			rowStyle = style.VeryLightGrayStyle
		case timeAgo < 9*statsboxEpoch:
			rowStyle = style.LightGrayStyle
		case timeAgo < 15*statsboxEpoch:
			rowStyle = style.GrayStyle
			printFaint = true
		default:
			printFaint = true
		}

		tokenHistory, ok := event.Other["fmtTokensHistory"].([]string)
		if !ok {
			gbl.Log.Warnf("could not get token history for event: %+v", event)

			continue
		}

		tokenInfo := strings.Join(tokenHistory, " | ")

		timeNow := rowStyle.Render(event.ReceivedAt.Format("15:04:05"))
		if event.IsOwnWallet {
			timeNow = collectionStyle.Render(event.ReceivedAt.Format("15:04:05"))
		}

		pricePerItem := price.NewPrice(event.Price.Wei())
		// TODO: fix this
		// if event.TransferredCollections.TotalTokens > 0 {
		if len(event.TransferredCollections) > 0 {
			// TODO: fix this (totaltokens)
			// pricePerItem = price.NewPrice(big.NewInt(0).Div(event.Price.Wei(), big.NewInt(event.TransferredCollections.TotalTokens)))
			pricePerItem = price.NewPrice(big.NewInt(0).Div(event.Price.Wei(), big.NewInt(int64(len(event.TransferredCollections)))))
		}

		historyLine := strings.Builder{}
		historyLine.WriteString(timeNow)
		historyLine.WriteString(" " + event.Typemoji)
		historyLine.WriteString(" " + rowStyle.Render(fmt.Sprintf("%6.3f", pricePerItem.Ether())))
		historyLine.WriteString(collectionStyle.Faint(printFaint).Render("Ξ"))
		// historyLine.WriteString(" " + event.TransferredCollections[0].CollectionName)
		historyLine.WriteString(" " + tokenInfo)

		if viper.GetBool("log.debug") {
			historyLine.WriteString(" " + fmt.Sprint(rowStyle.GetForeground()))
			historyLine.WriteString(" " + fmt.Sprint(rowStyle.GetFaint()))
		}

		eventsList = append(eventsList, listItem(historyLine.String()))
	}

	return eventsList
}

func (s *Stats) StartTicker(intervalPrintStats time.Duration, queueOutput chan string) {
	tickerPrintStats := time.NewTicker(time.Second * 7)

	gbl.Log.Infof("starting stats ticker (%s)", intervalPrintStats)

	time.Sleep(time.Until(time.Now().Truncate(intervalPrintStats).Add(intervalPrintStats)))

	tickerPrintStats.Reset(intervalPrintStats)

	go func() {
		for range tickerPrintStats.C {
			s.Print(queueOutput)
		}
	}()
}

func (s *Stats) runningForFullTimeframe() bool {
	return time.Since(internal.RunningSince) > viper.GetDuration("stats.timeframe")
}

func (s *Stats) getTimeframeMinutes() float64 {
	var timeframeMinutes float64
	if s.runningForFullTimeframe() {
		timeframeMinutes = viper.GetDuration("stats.timeframe").Minutes()
	} else {
		timeframeMinutes = time.Since(internal.RunningSince).Minutes()
	}

	return timeframeMinutes
}

func (s *Stats) highVolumeMint() {
	if viper.GetBool("show.mints") {
		log.Debug("👀 high volume mint | showing mints already active")

		return
	}

	checkInterval := viper.GetDuration("stats.high_volume.check_interval")
	minChecksBelowThreshold := viper.GetInt("stats.high_volume.min_checks_below_threshold")
	mintsTreshold := viper.GetInt("stats.high_volume.mints.threshold")

	mintsCount := s.eventslastTimeframe(degendb.Mint)
	mintsPerMin := int(float64(mintsCount) / s.getTimeframeMinutes())

	s.gb.Prf("👀 high volume mint (%d > %d /min| %d) | activating mints | check every: %.0fsec | min. checks below: %d", mintsPerMin, mintsTreshold, mintsCount, checkInterval.Seconds(), minChecksBelowThreshold)

	viper.Set("show.mints", true)

	// check if mintsPerMin is still above the threshold
	// otherwise deactivate displaying mints again
	checksBelow := 0
	for {
		// wait for next check
		time.Sleep(checkInterval)

		mintsCount := s.eventslastTimeframe(degendb.Mint)
		mintsPerMin := int(float64(mintsCount) / s.getTimeframeMinutes())

		log.Debugf("👀 -> high volume | last mints: %d <-> %d | checksBelow: %d <-> %d", mintsPerMin, mintsTreshold, checksBelow, minChecksBelowThreshold)

		if mintsPerMin < mintsTreshold {
			checksBelow++

			log.Debugf("👀 <- high volume | last mints below threshold %d < %d | checksBelow: %d <-> %d", mintsPerMin, mintsTreshold, checksBelow, minChecksBelowThreshold)

			if checksBelow >= minChecksBelowThreshold {
				s.gb.Prf("👀 high volume mint over, deactivating mints again")

				viper.Set("show.mints", false)

				return
			}
		}

		log.Debugf("👀 <- high volume | last mints: %d <-> %d | checksBelow: %d <-> %d", mintsPerMin, mintsTreshold, checksBelow, minChecksBelowThreshold)
	}
}

func GasTicker(gb *Gloomberg, gasTicker *time.Ticker, providerPool *provider.Pool, queueOutput chan string) {
	oldGasPrice := uint64(0)

	for range gasTicker.C {
		// gasNode := ethNodes.GetRandomLocalNode()
		gasLine := strings.Builder{}

		// if viper.GetBool("log.debug") {
		// 	gasLine.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#1A1A1A")).Render(fmt.Sprint(gasNode.Marker)))
		// }

		gasLine.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#111111")).Render("|"))
		gasLine.WriteString(style.GrayStyle.Copy().Faint(true).Render(time.Now().Format("15:04:05")))
		gasLine.WriteString(" " + style.DarkGrayStyle.Render("🧟"))

		gasLine.WriteString("   ")

		if gasInfo, err := providerPool.GetCurrentGasInfo(); err == nil && gasInfo != nil {
			// gas price
			if gasInfo.GasPriceWei.Cmp(big.NewInt(0)) > 0 {
				gasPriceGwei, _ := utils.WeiToGwei(gasInfo.GasPriceWei).Float64()
				gasPrice := uint64(math.Round(gasPriceGwei))

				atomic.StoreUint64(&gb.CurrentGasPriceGwei, gasPrice)

				if math.Abs(float64(gasPrice-oldGasPrice)) < 2.0 {
					continue
				}

				oldGasPrice = gasPrice

				// // tip / priority fee
				// var gasTip int
				// if gasInfo.GasTipWei.Cmp(big.NewInt(0)) > 0 {
				// 	gasTipGwei, _ := nodes.WeiToGwei(gasInfo.GasTipWei).Float64()
				// 	gasTip = int(math.Round(gasTipGwei))
				// }

				intro := style.DarkerGrayStyle.Render("~  ") + style.DarkGrayStyle.Render("gas") + style.DarkerGrayStyle.Render("  ~   ")
				outro := style.DarkerGrayStyle.Render("   ~   ~")
				divider := style.DarkerGrayStyle.Render("   ~   ~   ~   ~   ~   ~   ")

				formattedGas := style.GrayStyle.Render(fmt.Sprintf("%d", gasPrice)) + style.DarkGrayStyle.Render("gw")
				formattedGasAndTip := formattedGas

				// if gasTip > 0 {
				// 	formattedGasAndTip = formattedGas + "|" + style.GrayStyle.Render(fmt.Sprintf("%d", gasTip)) + style.DarkGrayStyle.Render("gw")
				// }

				gasLine.WriteString(intro + formattedGas + divider + formattedGasAndTip + divider + formattedGas + outro)
			}
		}

		queueOutput <- gasLine.String()
	}
}
