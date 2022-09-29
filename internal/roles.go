package internal

type Role string

func (r Role) String() string {
	return map[Role]string{
		ChainWatcher:          "ChainWatcher",
		OsStreamWatcher:       "OsStreamWatcher",
		OutputTerminal:        "OutputTerminal",
		OwnCollections:        "OwnCollections",
		OwnWalletWatcher:      "OwnWalletWatcher",
		StatsTicker:           "StatsTicker",
		TelegramBot:           "TelegramBot",
		TelegramNotifications: "TelegramNotifications",
		WalletWatcher:         "WalletWatcher",
		WsServer:              "WsServer",
	}[r]
}

const (
	ChainWatcher          Role = "chainWatcher"
	OsStreamWatcher       Role = "osStreamWatcher"
	OutputTerminal        Role = "outputTerminal"
	OwnCollections        Role = "ownCollections"
	OwnWalletWatcher      Role = "ownWalletWatcher"
	StatsTicker           Role = "statsTicker"
	TelegramBot           Role = "telegramBot"
	TelegramNotifications Role = "telegramNotifications"
	WalletWatcher         Role = "walletWatcher"
	WsServer              Role = "wsServer"
)

// type RoleMap map[Role]bool

type RoleMap struct {
	ChainWatcher          bool
	OsStreamWatcher       bool
	OutputTerminal        bool
	OwnCollections        bool
	OwnWalletWatcher      bool
	StatsTicker           bool
	TelegramBot           bool
	TelegramNotifications bool
	WalletWatcher         bool
	WsServer              bool
}

// func (r RoleMap) String() string {
// 	out := strings.Builder{}
// 	out.WriteString("chainWatcher: " + fmt.Sprint(r.ChainWatcher))
// 	out.WriteString("osStreamWatcher: " + fmt.Sprint(r.OsStreamWatcher))
// 	out.WriteString("wsServer: " + fmt.Sprint(r.WsServer))
// 	out.WriteString("tgBot: " + fmt.Sprint(r.TelegramBot))
// 	out.WriteString("TgNotifications: " + fmt.Sprint(r.TelegramBot))
// 	out.WriteString("outputTerminal: " + fmt.Sprint(r.OutputTerminal))

// 	return out.String()
// }
