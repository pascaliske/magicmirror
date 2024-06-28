package metrics

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// internal storage of metric updaters
var updaters = make(map[string]func())

func Middleware() echo.MiddlewareFunc {
	logger.Debug("Metrics endpoint enabled at %s", config.GetString("Metrics.Path"))

	// runtime metrics
	GoVersion := runtime.Version()
	Platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	// calculate uptime metrics
	uptime := NewUptimeMetrics()

	// metric defaults
	BuildInfo.WithLabelValues(version.Version, version.GitCommit, version.BuildTime, GoVersion, Platform).Set(1)
	UptimeSeconds.WithLabelValues().Set(0)
	ConfigReloadsTotal.WithLabelValues().Add(1)
	ConfigReloadsFailureTotal.WithLabelValues().Add(0)
	ConfigLastReloadSuccess.WithLabelValues().Set(uptime.GetStartup())
	ConfigLastReloadFailure.WithLabelValues().Set(0)

	// update config reload metrics
	config.OnChange("config-reload", func(success bool) {
		if success {
			ConfigReloadsTotal.WithLabelValues().Inc()
			ConfigLastReloadSuccess.WithLabelValues().SetToCurrentTime()
		} else {
			ConfigReloadsFailureTotal.WithLabelValues().Inc()
			ConfigLastReloadFailure.WithLabelValues().SetToCurrentTime()
		}
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// trigger metric updaters
			for _, callback := range updaters {
				callback()
			}

			// skip on metrics path
			if c.Path() == config.GetString("Metrics.Path") {
				return next(c)
			}

			// calculate http metrics
			err := next(c)
			http := NewHttpMetrics(c)

			// update metrics
			UptimeSeconds.WithLabelValues().Set(uptime.GetUptime())
			RequestsTotal.WithLabelValues(http.LabelsWithHost(err)...).Inc()
			RequestDuration.WithLabelValues(http.Labels(err)...).Observe(http.Duration())
			RequestSize.WithLabelValues(http.Labels(err)...).Observe(http.RequestSize())
			ResponseSize.WithLabelValues(http.Labels(err)...).Observe(http.ResponseSize())

			return err
		}
	}
}

func Handler() echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}

func OnRequest(id string, run func()) func() {
	// add callback to queue
	updaters[id] = func() {
		run()
	}

	// return unregister function
	return func() {
		delete(updaters, id)
	}
}
