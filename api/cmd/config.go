package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// flags
var output string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration file",
	Long:  "Manage the configuration file",

	Aliases: []string{"c"},
}

var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create configuration file",
	Long:  "Create a configuration file from defaults",

	Run: func(cmd *cobra.Command, args []string) {
		var defaultConfig config.Config

		// create a config instance with defaults
		c := config.CreateConfig()

		// unmarshal default values into variable
		err := c.Unmarshal(&defaultConfig)
		if err != nil {
			logger.Debug(err.Error())
			logger.Info("Could not create config with default values!")
			return
		}

		// marshal default config values as yaml
		data, err := yaml.Marshal(defaultConfig)
		if err != nil {
			logger.Debug(err.Error())
			logger.Error("Could not create config with default values!")
			return
		}

		// prepend yaml header
		data = append([]byte("---\n"), data...)

		// write to file if output flag is set
		if len(strings.Trim(output, " ")) > 0 {
			// normalize output file path
			output = filepath.Clean(strings.Trim(output, " "))

			// ensure absolute file path
			output, _ = filepath.Abs(output)

			// write generated config to file
			if err := os.WriteFile(output, data, 0640); err != nil {
				logger.Debug(err.Error())
				logger.Error("Could not write to given output file!")
				return
			}

			// inform user of successful config file creation
			logger.Info(fmt.Sprintf("Successfully created config file with default values: %s", output))
			return
		}

		// print generated config to stdout
		fmt.Print(string(data))
	},
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
	configCmd.AddCommand(configCreateCmd)
	configCmd.AddCommand(configValidateCmd)

	// flags
	configCreateCmd.Flags().StringVar(&output, "output", "", "Specify an output file instead of writing it to stdout")
}
