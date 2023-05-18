package stats

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync/atomic"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/nemo/wallet"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/lipgloss"
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

	interval time.Duration

	OwnEventsHistory []string
	EventHistory     []*totra.HistoryTokenTransaction

	gasTicker *time.Ticker

	salesVolume *big.Int
	sales       uint64
	mints       uint64

	NewLogs        uint64
	NewListings    uint64
	EventsToFormat uint64
	OutputLines    uint64

	DiscardedTransactions      uint64
	DiscardedTransfers         uint64
	DiscardedLowPrice          uint64
	DiscardedOtherERC          uint64
	DiscardedAlreadyKnownTX    uint64
	DiscardedUnknownCollection uint64
	DiscardedMints             uint64
}

func New(gasTicker *time.Ticker, wallets *wallet.Wallets, providerPool *provider.Pool, rdb rueidis.Client) *Stats {
	stats := &Stats{
		wallets:      wallets,
		providerPool: providerPool,
		rdb:          rdb,

		OwnEventsHistory: make([]string, viper.GetInt("stats.lines")),
		EventHistory:     make([]*totra.HistoryTokenTransaction, 0),

		gasTicker: gasTicker,

		interval: viper.GetDuration("ticker.statsbox"),
	}

	stats.Reset()

	return stats
}

func (s *Stats) salesPerMinute() float64 {
	return float64((s.sales * 60) / uint64(s.interval.Seconds()))
}

func (s *Stats) mintsPerMinute() float64 {
	return float64((s.mints * 60) / uint64(s.interval.Seconds()))
}

func (s *Stats) salesVolumePerMinute() float64 {
	ethVolume, _ := utils.WeiToEther(s.salesVolume).Float64()

	return (ethVolume * 60) / s.interval.Seconds()
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

func (s *Stats) AddSale(amountTokens int64, value *big.Int) float64 {
	s.salesVolume.Add(s.salesVolume, value)
	atomic.AddUint64(&s.sales, uint64(amountTokens))

	return float64((s.sales * 60) / uint64(s.interval.Seconds()))
}

func (s *Stats) AddMint(amountTokens int64) {
	atomic.AddUint64(&s.mints, uint64(amountTokens))
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
	// gbl.Log.Info("\n" + formattedStatsLists + "\n")

	s.Reset()
}

func (s *Stats) Reset() {
	gbl.Log.Debug("resetting statistics...")

	s.sales = 0
	s.mints = 0
	s.salesVolume = big.NewInt(0)
	s.DiscardedTransactions = 0
	s.DiscardedTransfers = 0
	s.DiscardedOtherERC = 0
	s.DiscardedAlreadyKnownTX = 0
	s.DiscardedUnknownCollection = 0
	s.DiscardedMints = 0
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
	} else if viper.IsSet("api_keys.etherscan") && viper.GetBool("stats.gas") {
		label := style.DarkGrayStyle.Render("  gas")
		value := style.LightGrayStyle.Render(fmt.Sprintf("%3d", external.GetEstimatedGasPrice()))

		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", label, value)), listItem("")}...)
	}

	//
	// per minute stats
	if volume := s.salesVolumePerMinute(); volume > 0 {
		volumeLabel := style.DarkGrayStyle.Render("Ξ /m")
		volumeValue := style.GrayStyle.Render(fmt.Sprintf("%6.2f", volume))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s%s", volumeValue, volumeLabel))}...)
	}

	if sales := s.salesPerMinute(); sales > 0 {
		salesLabel := style.DarkGrayStyle.Render("s/m")
		salesValue := style.GrayStyle.Render(fmt.Sprintf("%6d", uint(sales)))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", salesValue, salesLabel))}...)
	}

	if mints := s.mintsPerMinute(); mints > 0 {
		mintsLabel := style.DarkGrayStyle.Render("m/m")
		mintsValue := style.GrayStyle.Render(fmt.Sprintf("%6d", uint(mints)))
		firstColumn = append(firstColumn, []string{listItem(fmt.Sprintf("%s %s", mintsValue, mintsLabel))}...)
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
			var keyspaceHits, keyspaceMisses int

			for _, stat := range strings.Split(dbInfo, "\n") {
				if strings.HasPrefix(stat, "keyspace_hits:") {
					keyspaceHits, _ = fmt.Sscanf(stat, "keyspace_hits:%d", &keyspaceHits)
				}

				if strings.HasPrefix(stat, "keyspace_misses:") {
					keyspaceMisses, _ = fmt.Sscanf(stat, "keyspace_misses:%d", &keyspaceMisses)
				}
			}

			// calculate hitrate
			hitrate := float64(keyspaceHits) / float64(keyspaceHits+keyspaceMisses)

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
