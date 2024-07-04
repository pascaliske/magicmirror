package cmd

import (
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration file",
	Long:  "Manage the configuration file",

	Aliases: []string{"c"},
}

var configValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configuration file",
	Long:  "Validate the configuration file against it's schema",

	Run: func(cmd *cobra.Command, args []string) {
		// config seems valid
		logger.Info("Successfully validated config")
	},
}

func init() {
	cli.AddCommand(configCmd)

	// sub commands
	configCmd.AddCommand(configValidateCmd)
}
