package main

import (
	"os"

	"github.com/nuvosphere/nudex-voter/internal/config"
	"github.com/nuvosphere/nudex-voter/version"
	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "nudex-voter",
	Short: "nudex-voter is a tss(multi-party threshold signature scheme) signature server",
	Long:  `nudex-voter is a tss(multi-party threshold signature scheme) signature server`,
	Run: func(cmd *cobra.Command, args []string) {
		// print version
		version.PrintVersion(os.Stdout)
		config.InitConfig(configPath)
		app := NewApplication()
		app.Run()
	},
}

func init() {
	toolCmd.AddCommand(
		printPublicKeyCmd,
		printAddressCmd,
		printTssAddressCmd,
		printPeerIDCmd,
		printP2pFullAddressCmd,
		printECDSAPublicKeyCmd,
		printEDDSAPublicKeyCmd,
		printChainAddressCmd,
		printChainHotAddressCmd,
		printChainGasAddressCmd,
	)
	rootCmd.AddCommand(
		versionCmd,
		toolCmd,
	)

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "server config file")
}
