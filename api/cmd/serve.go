package cmd

import (
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server for API endpoints",
	Long:  "Start a HTTP server for providing the REST API endpoints",

	Aliases: []string{"start", "s"},

	PreRun: func(cmd *cobra.Command, args []string) {
		// print log level
		logger.Info("Log level set to %s", config.GetString("Log.Level"))

		// update log level on config change
		config.OnChangeSuccess("log-level", func() {
			logger.SetLevel(config.GetString("Log.Level"))
		})

		// watch config file
		config.Watch()
	},

	Run: func(cmd *cobra.Command, args []string) {
		server.SetupAndListen()
	},
}

func init() {
	cli.AddCommand(serveCmd)
}
