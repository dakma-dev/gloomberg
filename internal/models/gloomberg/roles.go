package gloomberg

type Role string

func (r Role) String() string {
	return map[Role]string{
		ChainWatcher:          "ChainWatcher",
		GloomClient:           "GloomClient",
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
	GloomClient           Role = "gloomClient"
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
	GloomClient           bool
	OsStreamWatcher       bool
	OutputTerminal        bool
	CollectionDB          bool
	OwnWalletWatcher      bool
	StatsTicker           bool
	TelegramBot           bool
	TelegramNotifications bool
	WalletWatcher         bool
	WsServer              bool
}
