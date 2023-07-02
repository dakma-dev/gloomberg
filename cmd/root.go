package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/benleb/gloomberg/cmd/flotscmd"
	"github.com/benleb/gloomberg/cmd/mintcmd"
	"github.com/benleb/gloomberg/cmd/oncecmd"
	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	ownWallets []string

	gb *gloomberg.Gloomberg
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gloomberg",
	Short: "A brief description of your application",
	Long:  style.GetSmallHeader(internal.GloombergVersion),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

//nolint:gochecknoinits
func init() {
	cobra.OnInitialize(initConfig)

	// subcommands
	rootCmd.AddCommand(flotscmd.FlotsCmd)
	rootCmd.AddCommand(mintcmd.MintCmd)
	rootCmd.AddCommand(oncecmd.OnceCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gloomberg.yaml)")

	// logging
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Show more output")
	_ = viper.BindPFlag("log.verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Show debug output")
	_ = viper.BindPFlag("log.debug", rootCmd.PersistentFlags().Lookup("debug"))

	// // rpc nodes
	// rootCmd.PersistentFlags().StringSliceVarP(&endpoints, "endpoints", "e", []string{}, "RPC endpoints")
	// _ = viper.BindPFlag("endpoints", rootCmd.Flags().Lookup("endpoints"))

	// // apis
	// rootCmd.PersistentFlags().StringVar(&apiKeyEtherscan, "etherscan", "", "Etherscan API Key")
	// _ = viper.BindPFlag("api_keys.etherscan", rootCmd.Flags().Lookup("etherscan"))
	// rootCmd.PersistentFlags().StringVar(&apiKeyMoralis, "moralis", "", "Moralis API Key")
	// _ = viper.BindPFlag("api_keys.moralis", rootCmd.Flags().Lookup("moralis"))
	// rootCmd.PersistentFlags().StringVar(&apiKeyOpensea, "opensea", "", "Opensea API Key")
	// _ = viper.BindPFlag("api_keys.opensea", rootCmd.Flags().Lookup("opensea"))

	// rootCmd.DebugFlags()
	// rootCmd.AddGroup(&cobra.Group{ID: "logging", Title: "logging"})
	// rootCmd.AddCommand(&cobraCommand{Use: "cmd1", GroupID: "group1", Run: emptyRun})
	// rootCmd.AddGroup(&cobra.Group{ID: "apikeys", Title: "api keys"})

	// // websockets server
	// rootCmd.PersistentFlags().Bool("server", false, "Start websockets server")
	// _ = viper.BindPFlag("server.enabled", rootCmd.Flags().Lookup("server"))
	// rootCmd.PersistentFlags().IP("host", net.IPv4(0, 0, 0, 0), "Websockets server port")
	// _ = viper.BindPFlag("server.host", rootCmd.Flags().Lookup("host"))
	// rootCmd.PersistentFlags().Uint16("port", 42069, "Websockets server port")
	// _ = viper.BindPFlag("server.port", rootCmd.Flags().Lookup("port"))

	// defaults

	// logging
	viper.SetDefault("log.log_file", "/tmp/gloomberg.log")

	// // api keys from nodes providers & other services
	// viper.SetDefault("api_keys", map[string]string{"alchemy": "", "infura": "", "moralis": "", "opensea": "", "etherscan": ""})

	// redis cache
	viper.SetDefault("redis.enabled", false)
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.database", 0)
	viper.SetDefault("redis.password", "")

	// ipfs
	// viper.SetDefault("ipfs.gateway", "https://ipfs.io/ipfs/")
	viper.SetDefault("ipfs.gateway", "https://cloudflare-ipfs.com/")

	// number of retries to resolve an ens name to an address or vice versa
	viper.SetDefault("ens.resolve_max_retries", 5)

	// collection/contract names
	viper.SetDefault("cache.names_ttl", 48*time.Hour)
	viper.SetDefault("cache.names_client_ttl", 1*time.Minute)

	// ens/wallet names
	viper.SetDefault("cache.ens_ttl", 48*time.Hour)

	// floor_ttl is only intended & suitable for caching purposes, not for buying decisions!
	viper.SetDefault("cache.floor_ttl", 10*time.Minute)
	viper.SetDefault("cache.salira_ttl", 1*time.Hour)
	viper.SetDefault("cache.slug_ttl", 3*24*time.Hour)
	viper.SetDefault("cache.notifications_lock_ttl", time.Millisecond*1337)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".gloomberg.yaml"
		viper.AddConfigPath(home)
		viper.SetConfigName(".gloomberg.yaml")
	}

	// config format
	viper.SetConfigType("yaml")

	// environment variables
	viper.SetEnvPrefix("GLOOMBERG")

	// read in environment variables that match
	viper.AutomaticEnv()

	// replace dots in env variables with underscores
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		//nolint:errorlint
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			fmt.Printf("config file error: %s - %s\n", viper.ConfigFileUsed(), err.Error())
		}
	}

	gbl.GetSugaredLogger()

	// // if command is not generate
	if rootCmd.CalledAs() != "generate" {
		gb = gloomberg.New()
	}

	fmt.Print(rootCmd.CalledAs())
}
