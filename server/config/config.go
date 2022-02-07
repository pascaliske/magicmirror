package config

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Port        int
	Metrics     struct {
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
	// fetch config file path
	input := flag.String("config", "", "")

	// parse flags
	flag.Parse()

	// read config file
	viper.AddConfigPath("/config")
	viper.AddConfigPath(filepath.Dir(*input))
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(*input), filepath.Ext(filepath.Base(*input))))
	viper.SetConfigType(strings.Replace(filepath.Ext(filepath.Base(*input)), ".", "", -1))

	// read environment variables
	viper.SetEnvPrefix("MM")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetTypeByDefaultValue(true)
	viper.AutomaticEnv()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// parse config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Config file found but not readable")
			return err
		}
	}

	return nil
}

func OnChange(run func()) {
	viper.OnConfigChange(func(e fsnotify.Event) {
		run()
	})
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
