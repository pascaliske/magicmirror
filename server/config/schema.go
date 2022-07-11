package config

import (
	"github.com/spf13/viper"
)

/**
 * Application environment. Can be "production" or "development".
 */
type Environment string

/**
 * Port for the server to listen on.
 */
type Port int

/**
 * Logging specific settings.
 */
type Log struct {
	Level string
}

/**
 * Metrics specific settings.
 */
type Metrics struct {
	Enabled bool
	Path    string
}

/**
 * Geo location settings. Pass-through to web app.
 */
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/**
 * News feed sources. Pass-through to web app.
 */
type Feeds []string

/**
 * API keys. Pass-through to web app.
 */
type ApiKeys struct {
	OpenWeather string `json:"openWeather"`
}

type Config struct {
	// general
	Environment Environment
	Port        Port
	Log         Log
	Metrics     Metrics

	// settings
	Location Location
	Feeds    Feeds

	// api keys
	ApiKeys ApiKeys
}

/**
 * Initialize default configuration values.
 */
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
