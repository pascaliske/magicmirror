package socket

import (
	"github.com/pascaliske/magicmirror/config"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Feed struct {
	Url string `json:"url"`
}

type ApiKeys struct {
	OpenWeather string `json:"openWeather"`
}

type Settings struct {
	Location Location `json:"location"`
	Feeds    []string `json:"feeds"`
	ApiKeys  ApiKeys  `json:"apiKeys"`
}

func BuildSettings() Settings {
	return Settings{
		Location: Location{
			Latitude:  config.GetFloat64("Location.Latitude"),
			Longitude: config.GetFloat64("Location.Longitude"),
		},
		Feeds: config.GetStringSlice("Feeds"),
		ApiKeys: ApiKeys{
			OpenWeather: config.GetString("ApiKeys.OpenWeather"),
		},
	}
}
