package metrics

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Middleware(Version string, GitCommit string, BuildTime string) echo.MiddlewareFunc {
	logger.Debug("Metrics endpoint enabled at %s", color.CyanString(config.GetString("Metrics.Path")))

	// runtime metrics
	GoVersion := runtime.Version()
	Platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	// calculate uptime metrics
	uptime := NewUptimeMetrics()

	// metric defaults
	BuildInfo.WithLabelValues(Version, GitCommit, BuildTime, GoVersion, Platform).Set(1)
	UptimeSeconds.WithLabelValues().Set(0)
	ConfigReloadsTotal.WithLabelValues().Add(1)
	ConfigReloadsFailureTotal.WithLabelValues().Add(0)
	ConfigLastReloadSuccess.WithLabelValues().Set(uptime.GetStartup())
	ConfigLastReloadFailure.WithLabelValues().Set(0)
	SocketConnections.WithLabelValues().Set(0)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
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

func Handler(server *echo.Echo) echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}
