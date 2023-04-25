package internal

import (
	"bufio"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

const (
	PubSubChannelSales = "sales"

	TopicSeaWatcher     = "seawatcher"
	TopicSeaWatcherMgmt = TopicSeaWatcher + "/mgmt"

	BlockTime = 12 * time.Second
)

var (
	GloombergVersion = "dev"

	WETHContractAddress           = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	BlurPoolTokenContractAddress  = common.HexToAddress("0x0000000000a39bb272e79075ade125fd351887ac")
	ENSContractAddress            = common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")
	ENSNameWrapperContractAddress = common.HexToAddress("0xd4416b13d2b3a9abae7acd5d6c2bbdbe25686401")

	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	ZeroHash    = common.Hash{}

	// BaseLogger is the logger used to print to the terminal without reporting caller or timestamp
	BaseLogger = log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
	})

	fileLogger = map[string]*log.Logger{}
)

// LoFi is the logger used to log to the log file with caller and timestamp reporting
func LoFi(filePath string) *log.Logger {
	if filePath == "" {
		filePath = viper.GetString("log.log_file")
	}

	if loFi, ok := fileLogger[filePath]; ok {
		return loFi
	}

	f, err := os.Create(filePath)
	if err != nil {
		BaseLogger.Error(err)
	}

	w := bufio.NewWriter(f)

	loFi := log.NewWithOptions(w, log.Options{
		TimeFormat:      time.DateTime,
		ReportCaller:    true,
		ReportTimestamp: true,
	})

	// defer f.Close()

	fileLogger[filePath] = loFi

	return loFi
}
