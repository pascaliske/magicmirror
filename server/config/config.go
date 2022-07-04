package config

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/pascaliske/magicmirror/logger"
	"github.com/spf13/viper"
)

/**
 * Parse a possible config file and watch for changes.
 */
func Parse() error {
	// define config file type
	viper.SetConfigType("yaml")

	// read config file from flag or default paths
	if found, file, dir := parseConfigFlag(); found {
		viper.SetConfigName(file)
		viper.AddConfigPath(dir)
	} else {
		viper.SetConfigName("config.yml")
		viper.AddConfigPath("/etc/magicmirror")
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
			return err
		}
	}

	// watch for config file changes
	watchConfig()

	return nil
}

/**
 * Parse a config flag if set and extract file and dir values.
 */
func parseConfigFlag() (bool, string, string) {
	// register config flag
	input := flag.String("config", "", "")

	// parse flags
	flag.Parse()

	// config flag not set
	if len(*input) == 0 {
		return false, "", ""
	}

	// ensure absolute directory
	file, _ := filepath.Abs(*input)

	// return file name and directory
	return true, filepath.Base(file), filepath.Dir(file)
}
