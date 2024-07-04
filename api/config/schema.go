package config

import (
	"github.com/spf13/viper"
)

/**
 * Application environment. Can be 'production' or 'development'.
 */
type Environment string

/**
 * Port for the server to listen on.
 */
type Port int

/**
 * Logging specific settings. Can be 'debug', 'info', 'warn', 'error', 'fatal'.
 */
type Log struct {
	Level string `validate:"required,oneof=debug info warn error fatal"`
}

/**
 * Metrics specific settings.
 */
type Metrics struct {
	Enabled bool
	Path    string `validate:"startswith=/,endsnotwith=/"`
}

/**
 * Geo location settings. Pass-through to web app.
 */
type Location struct {
	Latitude  float64 `json:"latitude" validate:"latitude"`
	Longitude float64 `json:"longitude" validate:"longitude"`
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
	// internal
	viper *viper.Viper `yaml:"-"`

	// general
	Environment Environment `yaml:"-" validate:"required,oneof=production development"`
	Port        Port        `validate:"required"`
	Log         Log
	Metrics     Metrics

	// settings
	Location Location
	Feeds    Feeds `validate:"unique,dive,url"`

	// api keys
	ApiKeys ApiKeys
}
