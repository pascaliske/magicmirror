package socket

import (
	"github.com/pascaliske/magicmirror/config"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ApiKeys struct {
	OpenWeather string `json:"openWeather"`
}

type Settings struct {
	Language string   `json:"language"`
	Units    string   `json:"units"`
	Location Location `json:"location"`
	ApiKeys  ApiKeys  `json:"apiKeys"`
}

func BuildSettings(cfg config.Config) Settings {
	return Settings{
		Language: cfg.Language,
		Units:    cfg.Units,
		Location: Location{
			Latitude:  cfg.Location.Latitude,
			Longitude: cfg.Location.Longitude,
		},
		ApiKeys: ApiKeys{
			OpenWeather: cfg.ApiKeys.OpenWeather,
		},
	}
}
