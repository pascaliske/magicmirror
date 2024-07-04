package config

func (c *Config) SetDefaults() {
	// general
	c.viper.SetDefault("Environment", "production")
	c.viper.SetDefault("Port", 9000)
	c.viper.SetDefault("Log.Level", "info")
	c.viper.SetDefault("Metrics.Enabled", true)
	c.viper.SetDefault("Metrics.Path", "/metrics")

	// settings
	c.viper.SetDefault("Location.Latitude", "")
	c.viper.SetDefault("Location.Longitude", "")
	c.viper.SetDefault("Feeds", []string{})

	// api keys
	c.viper.SetDefault("ApiKeys.OpenWeather", "")
}
