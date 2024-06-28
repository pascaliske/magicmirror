package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/version"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/**
 * Parse a possible config file and watch for changes.
 */
func ParseAndValidate(path string, checkMode bool, watchMode bool) {
	// define config file type
	viper.SetConfigType("yaml")

	// read config file from flag or default paths
	if found, file, dir := parseConfigPath(path); found {
		viper.SetConfigName(file)
		viper.AddConfigPath(dir)
	} else {
		viper.SetConfigName("config.yml")
		viper.AddConfigPath(fmt.Sprintf("/etc/%s", cases.Lower(language.English).String(version.Name)))
		viper.AddConfigPath(".")
	}

	// read environment variables
	viper.SetEnvPrefix("MM")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetTypeByDefaultValue(true)
	viper.AutomaticEnv()

	// parse config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Error("Config file found but not readable")
			os.Exit(2)
		}
	}

	// validate config
	if valid, err := validateConfig(); !valid {
		logger.Error(err.Error())
		os.Exit(2)
	}

	// exit if check mode is active
	if checkMode {
		logger.Info("Successfully validated config")
		os.Exit(0)
	}

	// watch for config file changes
	if watchMode {
		watchConfig()
	}
}

/**
 * Parse a config flag if set and extract file and dir values.
 */
func parseConfigPath(path string) (bool, string, string) {
	// config path is empty
	if len(path) == 0 {
		return false, "", ""
	}

	// ensure absolute directory
	file, _ := filepath.Abs(path)

	// return file name and directory
	return true, filepath.Base(file), filepath.Dir(file)
}
