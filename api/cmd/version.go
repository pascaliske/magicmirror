package cmd

import (
	"github.com/pascaliske/magicmirror/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  "Print version information of the MagicMirror instance and command-line interface",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintBannerWithVersion()
	},
}

func init() {
	cli.AddCommand(versionCmd)
}
