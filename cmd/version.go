package main

import (
	"os"

	"github.com/nuvosphere/nudex-voter/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of voter",
	Long:  `All software has versions. This is voter's`,
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintVersion(os.Stdout)
	},
}
