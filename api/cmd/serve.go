package cmd

import (
	"github.com/pascaliske/magicmirror/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server for API endpoints",
	Long:  "Start a HTTP server for providing the REST API endpoints",

	Aliases: []string{"start", "s"},

	Run: func(cmd *cobra.Command, args []string) {
		server.SetupAndListen()
	},
}

func init() {
	cli.AddCommand(serveCmd)
}
