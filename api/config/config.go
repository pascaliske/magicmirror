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

var c *Config

func init() {
	c = CreateConfig()
}

/**
 * Create a new viper instance with default values.
 */
func CreateConfig() *Config {
	c := new(Config)

	c.viper = viper.New()

	c.SetDefaults()

	return c
}

/**
 * Configure config file handling.
 */
func Setup(path string) {
	// define config file type
	c.viper.SetConfigType("yaml")

	// read config file from flag or default paths
	if found, file, dir := parseConfigPath(path); found {
		c.viper.SetConfigName(file)
		c.viper.AddConfigPath(dir)
	} else {
		c.viper.SetConfigName("config.yml")
		c.viper.AddConfigPath(fmt.Sprintf("/etc/%s", cases.Lower(language.English).String(version.Name)))
		c.viper.AddConfigPath(".")
	}

	// read environment variables
	c.viper.SetEnvPrefix("MM")
	c.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.viper.SetTypeByDefaultValue(true)
	c.viper.AutomaticEnv()
}

/**
 * Parse a possible config file.
 */
func ParseAndValidate() {
	// parse config
	if err := c.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Error("Config file found but not readable")
			os.Exit(2)
		}
	}

	// validate config
	if valid, err := c.validateConfig(); !valid {
		logger.Error(fmt.Sprintf("An error occurred: %s", err.Error()))
		os.Exit(2)
	}
}

/**
 * Unmarshal all values into the given variable.
 */
func (c *Config) Unmarshal(rawValue any) error {
	return c.viper.Unmarshal(&rawValue)
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
	file, err := filepath.Abs(path)

	if err != nil {
		return false, "", ""
	}

	// return file name and directory
	return true, filepath.Base(file), filepath.Dir(file)
}
