package config

import (
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/pascaliske/magicmirror/logger"
)

// internal storage of config file watchers
var watchers = make(map[string]func(bool))

/**
 * Watch config file for changes and notify all watchers.
 */
func Watch() {
	if file := c.viper.ConfigFileUsed(); len(file) > 0 {
		logger.Info("Watching for config file changes: %s", color.CyanString(file))
		c.viper.WatchConfig()
		c.viper.OnConfigChange(func(e fsnotify.Event) {
			// read-in config changes
			logger.Debug("Config file changed - %s", e.Op.String())
			success := c.viper.ReadInConfig() == nil

			if success {
				logger.Info("Config file reloaded successfully")
			} else {
				logger.Error("Config file reload failed")
			}

			// trigger config watchers
			for _, callback := range watchers {
				callback(success)
			}
		})
	}
}

/**
 * Register a config file watcher for all reloads.
 */
func OnChange(id string, run func(bool)) func() {
	// no config file used
	if len(c.viper.ConfigFileUsed()) == 0 {
		return func() {}
	}

	// add callback to queue
	watchers[id] = func(success bool) {
		run(success)
	}

	// return unregister function
	return func() {
		delete(watchers, id)
	}
}

/**
 * Register a config file watcher for only successful reloads.
 */
func OnChangeSuccess(id string, run func()) func() {
	return OnChange(id, func(success bool) {
		if success {
			run()
		}
	})
}

/**
 * Register a config file watcher for only failed reloads.
 */
func OnChangeError(id string, run func()) func() {
	return OnChange(id, func(success bool) {
		if !success {
			run()
		}
	})
}
