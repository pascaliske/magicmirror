package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Perform an health check against the API",
	Long:  "Perform an internal health check against the API endpoint",

	Aliases: []string{"status"},

	Run: func(cmd *cobra.Command, args []string) {
		// check health endpoint of application
		response, err := http.Get(fmt.Sprintf("http://localhost:%d/health", config.GetInt("Port")))

		// unhealthy
		if err != nil || response.StatusCode != http.StatusOK {
			logger.Debug(err.Error())
			logger.Error("Oh no! Something isn't alright here - the application seems to be unhealthy!")
			os.Exit(1)
		}

		// healthy
		logger.Info("Awesome! The application is in a healthy state!")
		os.Exit(0)
	},
}

func init() {
	cli.AddCommand(healthCmd)
}
