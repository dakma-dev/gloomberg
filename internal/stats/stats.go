package stats

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/totra"
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
	wallets      *wallet.Wallets
	providerPool *provider.Pool
	rdb          rueidis.Client

	interval  time.Duration
	timeframe time.Duration

	RecentEvents mapset.Set[*degendb.RecentEvent]

	OwnEventsHistory []string
	EventHistory     []*totra.HistoryTokenTransaction

	gasTicker *time.Ticker

	// salesVolume *big.Int
	// sales       uint64
	// mints       uint64

	NewLogs        uint64
	NewListings    uint64
	EventsToFormat uint64
	OutputLines    uint64
}

func New(gasTicker *time.Ticker, wallets *wallet.Wallets, providerPool *provider.Pool, rdb rueidis.Client) *Stats {
	stats := &Stats{
		wallets:      wallets,
		providerPool: providerPool,
		rdb:          rdb,

		RecentEvents: mapset.NewSet[*degendb.RecentEvent](),

		OwnEventsHistory: make([]string, viper.GetInt("stats.lines")),
		EventHistory:     make([]*totra.HistoryTokenTransaction, 0),

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

func (s *Stats) salesLastTimeframe() uint64 {
	count := uint64(0)

	for _, event := range s.RecentEvents.ToSlice() {
		if event.Type == degendb.Sale && time.Since(event.Timestamp) < s.timeframe {
			count += event.AmountTokens
		}
	}

	return count
}

func (s *Stats) mintsLastTimeframe() uint64 {
	count := uint64(0)

	for _, event := range s.RecentEvents.ToSlice() {
		if event.Type == degendb.Mint && time.Since(event.Timestamp) < s.timeframe {
			count += event.AmountTokens
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

		gbl.Log.Debugf("%s: %6.3fΞ total || %6.3f ETH | %6.3f WETH | %6.3f BlurPool", balanceTotalWei, balance.Account, utils.WeiToEther(balance.BalanceETH), utils.WeiToEther(balance.BalanceWETH), utils.WeiToEther(balance.BalanceBlurPool))

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

	if len(s.OwnEventsHistory) > 0 || len(s.EventHistory) > 0 {
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
			gasPrice := int(math.Ceil(gasPriceGwei))
			// gasTip, _ := nodes.WeiToGwei(gasInfo.GasTipWei).Uint64()

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

	// until runtime > timeframe, our data is not yet accurate -> darken the value
	valueStyle := style.GrayStyle
	if time.Since(internal.RunningSince) < viper.GetDuration("stats.timeframe") {
		valueStyle = style.DarkGrayStyle
	}

	// if volume := s.salesVolumePerMinute(); volume > 0 {
	if volume, _ := utils.WeiToEther(volumeWei).Float64(); volume > 0 {
		volumeLabel := style.DarkGrayStyle.Render("Ξ /m")

		// volumeValue := valueStyle.Render(fmt.Sprintf("%6.2f", volume))
		// firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s%s", volumeValue, volumeLabel))}...)

		volumePerMin := valueStyle.Render(fmt.Sprintf("%6.2f", volume/viper.GetDuration("stats.timeframe").Minutes()))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s%s", volumePerMin, volumeLabel))}...)
	}

	// if sales := s.salesPerMinute(); sales > 0 {
	if salesCount := s.salesLastTimeframe(); salesCount > 0 {
		salesLabel := style.DarkGrayStyle.Render("s/m")

		// salesValue := valueStyle.Render(fmt.Sprintf("%6d", salesCount))
		// firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", salesValue, salesLabel))}...)

		salesPerMin := valueStyle.Render(fmt.Sprintf("%6d", int(float64(salesCount)/viper.GetDuration("stats.timeframe").Minutes())))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", salesPerMin, salesLabel))}...)
	}

	// if mints := s.mintsPerMinute(); mints > 0 {
	if mintsCount := s.mintsLastTimeframe(); mintsCount > 0 {
		mintsLabel := style.DarkGrayStyle.Render("m/m")

		// mintsValue := valueStyle.Render(fmt.Sprintf("%6d", mintsCount))
		// firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", mintsValue, mintsLabel))}...)

		mintsPerMin := valueStyle.Render(fmt.Sprintf("%6d", int(float64(mintsCount)/viper.GetDuration("stats.timeframe").Minutes())))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", mintsPerMin, mintsLabel))}...)
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
			hitrate := float64(keyspaceHits) / float64(keyspaceHits+keyspaceMisses)
			log.Debugf("keyspace_hits: %+v | keyspace_misses: %+v | hitrate: %+v", keyspaceHits, keyspaceMisses, hitrate)

			// cache size
			namesLabel := style.DarkGrayStyle.Render("  cache")
			namesValue := style.GrayStyle.Render(fmt.Sprintf("%9d", dbSize))

			// salesLabel := style.DarkGrayStyle.Render("s-cache")
			// salesValue := style.GrayStyle.Render(fmt.Sprintf("%9d", rdb.XLen(context.Background(), "sales").Val()))

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

	if len(s.EventHistory) == 0 {
		gbl.Log.Debugf("no events to show")

		return eventsList
	}

	// cleanup (maybe replace this by not inserting events that are not shown anyways)
	historyEvents := make([]*totra.HistoryTokenTransaction, 0)

	for idx, event := range s.EventHistory {
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

		if len(event.FmtTokensTransferred) == 0 {
			continue
		}

		collectionStyle := lipgloss.NewStyle().Foreground(event.Collection.Colors.Primary)

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

		tokenInfo := event.FmtTokensTransferred[0] // strings.Join(event.FmtTokensTransferred, " | ")

		isOwnWallet := s.wallets.ContainsAddressFromSlice(event.TokenTransaction.GetNFTSenderAndReceiverAddresses()) != internal.ZeroAddress

		timeNow := rowStyle.Render(event.ReceivedAt.Format("15:04:05"))
		if isOwnWallet {
			timeNow = collectionStyle.Render(event.ReceivedAt.Format("15:04:05"))
		}

		pricePerItem := price.NewPrice(event.AmountPaid)
		if event.TokenTransaction.TotalTokens > 0 {
			pricePerItem = price.NewPrice(big.NewInt(0).Div(event.AmountPaid, big.NewInt(event.TokenTransaction.TotalTokens)))
		}

		historyLine := strings.Builder{}
		historyLine.WriteString(timeNow)
		historyLine.WriteString(" " + event.TokenTransaction.Action.Icon())
		historyLine.WriteString(" " + rowStyle.Render(fmt.Sprintf("%6.3f", pricePerItem.Ether())))
		historyLine.WriteString(collectionStyle.Faint(printFaint).Render("Ξ"))
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
