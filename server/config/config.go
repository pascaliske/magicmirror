package config

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/spf13/viper"
)

var callbacks = make(map[string]func(bool))

func Parse() error {
	// define config file type
	viper.SetConfigType("yaml")

	// read config file from flag or default paths
	if found, file, dir := parseConfigFlag(); found {
		viper.SetConfigName(file)
		viper.AddConfigPath(dir)
	} else {
		viper.SetConfigName("config.yml")
		viper.AddConfigPath("/etc/magicmirror")
		viper.AddConfigPath(".")
	}

	// read environment variables
	viper.SetEnvPrefix("MM")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetTypeByDefaultValue(true)
	viper.AutomaticEnv()

	// parse config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Error("Config file found but not readable")
			return err
		}
	}

	// watch for config file changes
	if file := viper.ConfigFileUsed(); len(file) > 0 {
		logger.Info("Watching for config file changes: %s", color.CyanString(file))
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			logger.Info("Config file changed")

			// check if change is valid
			valid := viper.ReadInConfig() == nil

			for _, callback := range callbacks {
				callback(valid)
			}
		})
	}

	return nil
}

func parseConfigFlag() (bool, string, string) {
	// register config flag
	input := flag.String("config", "", "")

	// parse flags
	flag.Parse()

	// config flag not set
	if len(*input) == 0 {
		return false, "", ""
	}

	// ensure absolute directory
	file, _ := filepath.Abs(*input)

	// return file name and directory
	return true, filepath.Base(file), filepath.Dir(file)
}

func OnChange(id string, run func(bool)) func() {
	// no config file used
	if len(viper.ConfigFileUsed()) == 0 {
		return func() {}
	}

	// add callback to queue
	callbacks[id] = func(success bool) {
		run(success)
	}

	// return unregister function
	return func() {
		delete(callbacks, id)
	}
}

func OnChangeSuccess(id string, run func()) func() {
	return OnChange(id, func(success bool) {
		if success {
			run()
		}
	})
}

func OnChangeError(id string, run func()) func() {
	return OnChange(id, func(success bool) {
		if !success {
			run()
		}
	})
}
