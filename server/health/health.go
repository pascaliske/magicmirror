package health

import (
	"github.com/fatih/color"
	"github.com/hellofresh/health-go/v4"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/logger"
)

func Handler(server *echo.Echo) echo.HandlerFunc {
	logger.Debug("Health endpoint ready at %s", color.CyanString("/health"))

	h, err := health.New()

	if err != nil {
		server.Logger.Fatal(err)
	}

	return echo.WrapHandler(h.Handler())
}
