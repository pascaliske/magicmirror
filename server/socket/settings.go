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

func BuildSettings(cfg config.Config) Settings {
	return Settings{
		Location: Location{
			Latitude:  cfg.Location.Latitude,
			Longitude: cfg.Location.Longitude,
		},
		Feeds: cfg.Feeds,
		ApiKeys: ApiKeys{
			OpenWeather: cfg.ApiKeys.OpenWeather,
		},
	}
}
