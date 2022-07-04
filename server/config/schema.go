package config

import (
	"github.com/spf13/viper"
)

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
	viper.SetDefault("Log.Level", "info")
	viper.SetDefault("Metrics.Enabled", true)
	viper.SetDefault("Metrics.Path", "/metrics")

	// settings
	viper.SetDefault("Location.Latitude", "")
	viper.SetDefault("Location.Longitude", "")
	viper.SetDefault("Feeds", []string{})

	// api keys
	viper.SetDefault("ApiKeys.OpenWeather", "")
}
