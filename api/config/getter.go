package config

func GetBool(key string) bool {
	return c.viper.GetBool(key)
}

func GetString(key string) string {
	return c.viper.GetString(key)
}

func GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

func GetInt(key string) int {
	return c.viper.GetInt(key)
}

func GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}
