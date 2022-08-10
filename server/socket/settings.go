package socket

import (
	"github.com/pascaliske/magicmirror/config"
)

type Settings struct {
	Location config.Location `json:"location"`
	Feeds    config.Feeds    `json:"feeds"`
	ApiKeys  config.ApiKeys  `json:"apiKeys"`
}

func (client Client) BuildSettings() Settings {
	return Settings{
		Location: config.Location{
			Latitude:  config.GetFloat64("Location.Latitude"),
			Longitude: config.GetFloat64("Location.Longitude"),
		},
		Feeds: config.GetStringSlice("Feeds"),
		ApiKeys: config.ApiKeys{
			OpenWeather: config.GetString("ApiKeys.OpenWeather"),
		},
	}
}
