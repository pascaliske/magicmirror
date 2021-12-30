package metrics

import (
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	// setup metrics
}

func Handler(cfg config.Config, server *echo.Echo) echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}
