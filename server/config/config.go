package config

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/spf13/viper"
)

var callbacks = make(map[string]func())

type Config struct {
	// general
	Environment string
	Port        int
	Log         struct {
		Level string
	}
	Metrics struct {
		Enabled bool
		Path    string
	}

	// settings
	Location struct {
		Latitude  float64
		Longitude float64
	}
	Feeds []string

	// api keys
	ApiKeys struct {
		OpenWeather string
	}
}

func init() {
	// general
	viper.SetDefault("Environment", "production")
	viper.SetDefault("Port", 9000)
	viper.SetDefault("Metrics.Enabled", true)
	viper.SetDefault("Metrics.Path", "/metrics")

	// settings
	viper.SetDefault("Location.Latitude", "")
	viper.SetDefault("Location.Longitude", "")
	viper.SetDefault("Feeds", []string{})

	// api keys
	viper.SetDefault("ApiKeys.OpenWeather", "")
}

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
	if file := viper.ConfigFileUsed(); len(file) > 0 {
		logger.Debug("Watching for config file changes: %s", color.CyanString(file))
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			logger.Info("Config file changed")

			for _, callback := range callbacks {
				callback()
			}
		})
	}

	return nil
}

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

func OnChange(id string, run func()) func() {
	// no config file used
	if len(viper.ConfigFileUsed()) == 0 {
		return func() {}
	}

	// add callback to queue
	callbacks[id] = run

	// return unregister function
	return func() {
		delete(callbacks, id)
	}
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}
