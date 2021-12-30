package health

import (
	"github.com/hellofresh/health-go/v4"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
)

func Handler(cfg config.Config, server *echo.Echo) echo.HandlerFunc {
	h, err := health.New()

	if err != nil {
		server.Logger.Fatal(err)
	}

	return echo.WrapHandler(h.Handler())
}
