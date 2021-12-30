package config

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

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
	Units    string
	Language string
	Location struct {
		Latitude  float64
		Longitude float64
	}
}

func init() {
	viper.SetDefault("Environment", "production")
	viper.SetDefault("Port", 9000)
	viper.SetDefault("Metrics.Enabled", true)
	viper.SetDefault("Metrics.Path", "/metrics")

	// settings
	viper.SetDefault("Units", "")
	viper.SetDefault("Language", "")
	viper.SetDefault("Location.Latitude", "")
	viper.SetDefault("Location.Longitude", "")
}

func Parse() (config Config, err error) {
	// fetch config file path
	input := flag.String("config", "", "")

	// parse flags
	flag.Parse()

	// read config file
	viper.AddConfigPath(filepath.Dir(*input))
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(*input), filepath.Ext(filepath.Base(*input))))
	viper.SetConfigType(strings.Replace(filepath.Ext(filepath.Base(*input)), ".", "", -1))

	// read environment variables
	viper.SetEnvPrefix("MM")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// parse config
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Config file found but not readable")
			return
		}
	}

	// return parsed values
	err = viper.Unmarshal(&config)
	return
}
