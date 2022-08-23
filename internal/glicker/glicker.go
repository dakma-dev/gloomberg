package glicker

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
	"github.com/benleb/gloomberg/internal/cache"
	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/etherscan"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/subscriptions"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

var (
	ctx context.Context

	columnWidth = 32
	// subtle      = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	// highlight := lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	// special := lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}.
	listStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(style.Subtle).
			MarginRight(0).
		// Height(8).
		Width(columnWidth + 1)

	// listHeader = lipgloss.NewStyle().
	// 		BorderStyle(lipgloss.NormalBorder()).
	// 		BorderBottom(true).
	// 		BorderForeground(subtle).
	// 		MarginRight(2).
	// 		Render

	// checkMark := lipgloss.NewStyle().SetString("✓").
	// 	Foreground(special).
	// 	PaddingRight(1).
	// 	String()

	// listDone := func(s string) string {
	// 	return checkMark + lipgloss.NewStyle().
	// 		Strikethrough(true).
	// 		Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
	// 		Render(s)
	// }.

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render

	StatsTicker *Stats
)

type Stats struct {
	wallets *models.Wallets

	numCollections uint
	interval       time.Duration

	OwnEventsHistory []string
	EventHistory     []*collections.Event

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

func New(context context.Context, wallets *models.Wallets, numCollections int) *Stats {
	ctx = context

	stats := &Stats{
		wallets: wallets,

		OwnEventsHistory: make([]string, viper.GetInt("stats.lines")),
		EventHistory:     make([]*collections.Event, 0),

		numCollections: uint(numCollections),
		interval:       viper.GetDuration("stats.interval"),
	}

	stats.Reset()

	StatsTicker = stats

	return stats
}

func (s *Stats) salesPerMinute() float64 {
	return float64((s.sales * 60) / uint64(s.interval.Seconds()))
}

// func (s *Stats) mintsPerMinute() float64 {
// 	var mints uint64

// 	if viper.GetBool("show.mints") {
// 		mints = s.mints
// 	} else {
// 		mints = s.DiscardedMints
// 	}

// 	return float64((mints * 60) / uint64(s.interval.Seconds()))
// }

func (s *Stats) salesVolumePerMinute() float64 {
	ethVolume, _ := subscriptions.WeiToEther(s.salesVolume).Float64()

	return float64((ethVolume * 60) / s.interval.Seconds())
}

func (s *Stats) processedLogs() uint64 {
	return ((s.sales +
		s.DiscardedTransactions +
		s.DiscardedTransfers +
		s.DiscardedOtherERC +
		s.DiscardedAlreadyKnownTX +
		s.DiscardedUnknownCollection +
		s.DiscardedMints) * 60) / uint64(s.interval.Seconds())
}

func (s *Stats) UpdateBalances() {
	gbl.Log.Debugf("updating wallet balances...")

	balances := etherscan.GetBalances(s.wallets)
	if balances == nil {
		gbl.Log.Error("❌ error while fetching wallet balances")
		return
	}

	if viper.GetBool("log.debug") {
		for _, balance := range balances {
			gbl.Log.Debugf("UpdateBalances| %+v\n", balance)
		}
	}

	for _, balance := range balances {
		walletAddress := common.HexToAddress(balance.Account)

		balanceTotalWei := big.NewInt(0).Add(balance.BalanceETH, balance.BalanceWETH)

		(*s.wallets)[walletAddress].BalanceBefore = (*s.wallets)[walletAddress].Balance
		(*s.wallets)[walletAddress].Balance = balanceTotalWei

		trendIndicator := style.CreateTrendIndicator(
			float64((*s.wallets)[walletAddress].BalanceBefore.Int64()),
			float64((*s.wallets)[walletAddress].Balance.Int64()),
		)

		(*s.wallets)[walletAddress].BalanceTrend = trendIndicator

		gbl.Log.Debugf("  %s balance: %s %6.3f", balance.Account, trendIndicator, subscriptions.WeiToEther((*s.wallets)[walletAddress].Balance))
	}
}

func (s *Stats) AddSale(value *big.Int) float64 {
	s.salesVolume.Add(s.salesVolume, value)
	atomic.AddUint64(&s.sales, 1)

	return float64((s.sales * 60) / uint64(s.interval.Seconds()))
}

func (s *Stats) AddMint() {
	atomic.AddUint64(&s.mints, 1)
}

func (s *Stats) Print() {
	var (
		formattedStatsLists string

		statsLists []string
	)

	if viper.GetBool("stats.balances") {
		s.UpdateBalances()
	}

	// new
	statsLists = []string{}
	statsLists = append(statsLists, s.getPrimaryStatsLists()...)

	if viper.GetBool("stats.extended") {
		statsLists = append(statsLists, s.getDiscardsStatsList()...)
		statsLists = append(statsLists, s.getChannelStatsList()...)
	}

	maxWalletNameLength := 0.0
	for _, wallet := range *s.wallets {
		maxWalletNameLength = math.Max(maxWalletNameLength, float64(len(wallet.Name)))
	}

	if walletBalancesList := s.getWalletStatsList(int(maxWalletNameLength)); len(walletBalancesList) > 0 {
		walletColumnWidth := int(math.Min(float64(columnWidth), maxWalletNameLength)) + 12
		sizedWalletList := listStyle.Copy().Width(walletColumnWidth).Render
		statsLists = append(statsLists, sizedWalletList(lipgloss.JoinVertical(lipgloss.Left, walletBalancesList...)))
	}

	if len(s.OwnEventsHistory) > 0 || len(s.EventHistory) > 0 {
		eventsList := listStyle.Copy().UnsetWidth().UnsetBorderRight().PaddingLeft(0).Render
		statsLists = append(statsLists, eventsList(lipgloss.JoinVertical(lipgloss.Left, s.getOwnEventsHistoryList()...)))
	}

	formattedStatsLists = lipgloss.JoinHorizontal(lipgloss.Top, statsLists...)

	fmt.Println("")
	fmt.Println(formattedStatsLists)
	fmt.Println("")

	s.Reset()
}

func (s *Stats) Reset() {
	gbl.Log.Infof("resetting statistics...")

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

	if viper.IsSet("api_keys.etherscan") && viper.GetBool("stats.gas") {
		firstColumn = append(firstColumn, []string{
			// listItem(lipgloss.NewStyle().Align(lipgloss.Right).Render(fmt.Sprintf(
			// 	" ~%s  %s",
			// 	lipgloss.NewStyle().Bold(true).Render(GetEstimatedGasPrice().String()),
			// 	lipgloss.NewStyle().Faint(true).Render("gas"),
			// ))),
			listItem(formatCounter(style.DarkWhiteStyle.Copy().Align(lipgloss.Right).Render(fmt.Sprintf("%4d", etherscan.GetEstimatedGasPrice())), "𛱟 gas")),
			listItem(""),
		}...)
	}

	firstColumn = append(firstColumn, []string{
		// listItem(FormatCounter(lipgloss.NewStyle().Align(lipgloss.Right).Render(fmt.Sprintf("%4d", GetEstimatedGasPrice())), "𛱟 gas")),
		listItem(formatCounter(fmt.Sprintf("%.1f", s.salesVolumePerMinute()), "Ξ vol/min")),
		listItem(formatCounter(fmt.Sprintf("%4d", uint(s.salesPerMinute())), " sales/min")),
		// listItem(formatCounter(fmt.Sprintf("%4d", uint(s.mintsPerMinute())), "  mints/min")),
	}...)

	// second column
	var secondcolumn []string

	if minPrice := viper.GetFloat64("show.min_price"); minPrice > 0.0 {
		secondcolumn = append(secondcolumn, []string{
			listItem(formatCounter(fmt.Sprintf("%4.2f", minPrice), "Ξ min price")),
		}...)
	}

	if viper.GetBool("redis.enabled") {
		if rdb := cache.New(ctx).GetRDB(); rdb != nil {
			secondcolumn = append(secondcolumn, []string{
				listItem(formatCounter(rdb.DBSize(context.Background()).Val(), " in n-cache")),
				listItem(formatCounter(rdb.XLen(context.Background(), "sales").Val(), " in s-cache")),
			}...)
		}
	}

	if processedLogs := s.processedLogs(); processedLogs > 0 {
		secondcolumn = append(secondcolumn, []string{
			listItem(formatCounter(s.processedLogs(), " logs/min")),
		}...)
	}

	statsOutput := []string{listStyle.Copy().Width(17).Render(lipgloss.JoinVertical(lipgloss.Left, firstColumn...))}

	if len(secondcolumn) > 0 {
		statsOutput = append(statsOutput, listStyle.Copy().Width(20).Render(lipgloss.JoinVertical(lipgloss.Left, secondcolumn...)))
	}

	return statsOutput
}

func (s *Stats) getWalletStatsList(maxWalletNameLength int) []string {
	wallets := s.wallets.GetAll()
	sort.Sort(sort.Reverse(wallets))

	numberOfWalletsToShow := int(math.Min(float64(viper.GetInt("stats.lines")), float64(len(wallets))))

	walletsList := make([]string, 0)

	for _, wallet := range wallets[:numberOfWalletsToShow] {
		walletStats := listItem(fmt.Sprintf("%s %s %5.2f%s", wallet.ColoredName(maxWalletNameLength), wallet.BalanceTrend, subscriptions.WeiToEther(wallet.Balance), style.GrayStyle.Render("Ξ")))
		walletsList = append(walletsList, walletStats)
	}

	return walletsList
}

func (s *Stats) getDiscardsStatsList() []string {
	var discardStats []string

	discardStats = append(discardStats,
		listItem(fmt.Sprint(formatCounter(s.DiscardedAlreadyKnownTX, " known"), "  |  ", formatCounter(s.DiscardedUnknownCollection, " unknown"))),
		listItem(fmt.Sprint(formatCounter(s.DiscardedTransfers, "tf"), "  |  ", formatCounter(s.DiscardedTransactions, "tx"))),
		listItem(fmt.Sprint(formatCounter(s.DiscardedLowPrice, "low"), "  |  ", formatCounter(s.DiscardedMints, "mt"))),
		listItem(formatCounter(s.DiscardedOtherERC, " non-erc721/1155")),
	)

	return []string{
		listStyle.Copy().Width(26).Render(lipgloss.JoinVertical(lipgloss.Left, discardStats...)),
	}
}

func (s *Stats) getChannelStatsList() []string {
	var channelStats []string

	channelStats = append(channelStats,
		// listItem(fmt.Sprintf("%s %s%s", lipgloss.NewStyle().Faint(true).Width(16).Render(fmt.Sprintf("%s", "newLogs")), fmt.Sprintf("%d", s.NewLogs), GrayStyle.Render("msgs"))),
		// // listItem(fmt.Sprintf("%s %d%s", lipgloss.NewStyle().Faint(true).Width(16).Render(fmt.Sprintf("%s", "newListings")), len(NewListings), GrayStyle.Render("msgs"))),
		// listItem(fmt.Sprintf("%s %s%s", lipgloss.NewStyle().Faint(true).Width(16).Render(fmt.Sprintf("%s", "salesToFormat")), fmt.Sprintf("%d", s.EventsToFormat), GrayStyle.Render("msgs"))),
		// listItem(fmt.Sprintf("%s %s%s", lipgloss.NewStyle().Faint(true).Width(16).Render(fmt.Sprintf("%s", "outputLines")), fmt.Sprintf("%d", s.OutputLines), GrayStyle.Render("msgs"))),

		// listItem(FormatCounter(scheduler.NewLogs, " newLogs")),
		listItem(formatCounter(s.EventsToFormat, " salesToFormat")),
		listItem(formatCounter(s.OutputLines, " outputLines")),
		listItem(formatCounter(s.NewListings, " newListings")),
	)

	return []string{
		listStyle.Copy().Width(20).Render(lipgloss.JoinVertical(lipgloss.Left, channelStats...)),
	}
}

func (s *Stats) getOwnEventsHistoryList() []string {
	var eventsList []string

	// if numberOfOwnEvents := len(s.OwnEventsHistory); numberOfOwnEvents > 0 {
	// 	numberOfShownEvents := int(math.Min(float64(viper.GetInt("stats.lines")), float64(numberOfOwnEvents)))
	// 	ownEvents := s.OwnEventsHistory[numberOfOwnEvents-numberOfShownEvents:]
	// 	sort.Slice(ownEvents, func(i, j int) bool { return i < j })

	// 	for _, event := range ownEvents {
	// 		eventsList = append(eventsList, listItem(event))
	// 	}
	// }

	if numberOfOwnEvents := len(s.EventHistory); numberOfOwnEvents > 0 {
		gbl.Log.Debugf("numberOfOwnEvents: %d | %d\n", numberOfOwnEvents, len(s.EventHistory))

		numberOfShownEvents := int(math.Min(float64(viper.GetInt("stats.lines")), float64(numberOfOwnEvents)))
		ownEvents := s.EventHistory[numberOfOwnEvents-numberOfShownEvents:]
		sort.Slice(ownEvents, func(i, j int) bool { return i < j })

		gbl.Log.Debugf("ownEvents: %d", len(ownEvents))

		for idx, event := range ownEvents {
			if event == nil {
				gbl.Log.Debugf("event is nil: %d\n", idx)

				continue
			}

			gbl.Log.Debugf("%d | event: %+v\n", idx, event)

			// rCol := fmt.Sprintf("#%x", 0x111111*((idx+1)*3))

			// rCol := fmt.Sprintf("#%x", 0x111111*((idx+3)*2))

			// toSub := len(ownEvents) - idx
			// r, g, b := 0x11, 0x11, 0x11
			// darkenFactor := 1.4
			// rCol := fmt.Sprintf(
			// 	"#%x%x%x",
			// 	0xFF-(int(float64(toSub)*float64(r)*darkenFactor)),
			// 	0xFF-(int(float64(toSub)*float64(g)*darkenFactor)),
			// 	0xFF-(int(float64(toSub)*float64(b)*darkenFactor)),
			// )

			var rowStyle lipgloss.Style

			collectionStyle := lipgloss.NewStyle().Foreground(event.CollectionColor)

			timeAgo := time.Since(event.Time)
			glickerEpoch := viper.GetDuration("stats.interval")

			printFaint := false

			switch {
			case timeAgo < glickerEpoch:
				rowStyle = style.DarkWhiteStyle
			case timeAgo < 2*glickerEpoch:
				rowStyle = style.VeryLightGrayStyle
			case timeAgo < 4*glickerEpoch:
				rowStyle = style.LightGrayStyle
			case timeAgo < 8*glickerEpoch:
				rowStyle = style.GrayStyle
				printFaint = true
			default:
				rowStyle = style.DarkGrayStyle
				printFaint = true
			}

			var tokenInfo string
			if event.TxItemCount > 1 {
				tokenInfo = fmt.Sprintf("%s %s", rowStyle.Render(fmt.Sprintf("%dx", event.TxItemCount)), collectionStyle.Faint(printFaint).Render(event.Collection.Name))
			} else {
				tokenInfo = internal.FormatTokenInfo(event.TokenID, event.Collection, printFaint, true)
			}

			timeNow := rowStyle.Render(event.Time.Format("15:04:05"))

			historyLine := strings.Builder{}
			historyLine.WriteString(timeNow)
			// historyLine.WriteString(" " + timeAgo.Truncate(time.Second).String())
			historyLine.WriteString(" " + event.EventType.Icon())
			historyLine.WriteString(" " + rowStyle.Render(fmt.Sprintf("%6.3f", subscriptions.WeiToEther(event.PricePerItem))))
			historyLine.WriteString(collectionStyle.Faint(printFaint).Render("Ξ"))
			historyLine.WriteString(" " + tokenInfo)

			if viper.GetBool("log.debug") {
				historyLine.WriteString(" " + fmt.Sprint(rowStyle.GetForeground()))
				historyLine.WriteString(" " + fmt.Sprint(rowStyle.GetFaint()))
			}

			eventsList = append(eventsList, listItem(historyLine.String()))
		}
	}

	return eventsList
}

func (s *Stats) StartTicker(intervalPrintStats time.Duration) {
	intervalPrintStats = viper.GetDuration("stats.interval")
	tickerPrintStats := time.NewTicker(time.Second * 7)

	gbl.Log.Infof("starting stats ticker - updating every %s", intervalPrintStats)

	go func() {
		time.Sleep(time.Until(time.Now().Truncate(intervalPrintStats).Add(intervalPrintStats)))

		tickerPrintStats.Reset(intervalPrintStats)

		for range tickerPrintStats.C {
			// if err := viper.WriteConfigAs(viper.GetString("config.backup_file")); err != nil {
			// 	gbl.Log.Errorf("writing config backup to %s failed: %s", viper.GetString("config.backup_file"), err)
			// }
			s.Print()
		}
	}()
}

func formatCounter(value any, label string) string {
	var valueString string

	switch val := value.(type) {
	case int, int64, uint, uint64:
		valueString = fmt.Sprint(val)

	case string:
		valueString = val
	}

	counterStyle := style.GrayStyle.Copy()
	counterName := counterStyle.Faint(true).Render(label)
	counterValue := counterStyle.Faint(false).Render(valueString)

	return fmt.Sprintf("%s%s", counterValue, counterName)
}
