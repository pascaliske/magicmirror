package cmd

import (
	"fmt"

	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// flags
var configPath string

var cli = &cobra.Command{
	Use:  cases.Lower(language.English).String(version.Name),
	Long: fmt.Sprintf("Manage the %s instance from the command-line", version.Name),

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// print version banner only on serve command
		if cmd.CalledAs() == serveCmd.Name() {
			version.PrintBannerWithVersion()
		}

		// parse and validate config file
		config.ParseAndValidate(configPath, false, cmd.CalledAs() == serveCmd.Name())

		// configure log level
		logger.SetLevel(config.GetString("Log.Level"))
		config.OnChangeSuccess("log-level", func() {
			logger.SetLevel(config.GetString("Log.Level"))
		})
	},
}

func init() {
	// general settings
	cli.CompletionOptions.HiddenDefaultCmd = true

	// config flag
	cli.PersistentFlags().StringVar(&configPath, "config", "", "Path to configuration file")
	if err := viper.BindPFlag("config", cli.PersistentFlags().Lookup("config")); err != nil {
		logger.Error(err.Error())
	}
}

func Execute() error {
	return cli.Execute()
}
