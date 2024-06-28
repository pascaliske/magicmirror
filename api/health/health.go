package health

import (
	"fmt"

	"github.com/hellofresh/health-go/v5"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/version"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Handler() echo.HandlerFunc {
	logger.Debug("Health endpoint ready at %s", "/health")

	h, _ := health.New(health.WithComponent(health.Component{
		Name:    cases.Lower(language.English).String(version.Name),
		Version: fmt.Sprintf("%s @ %s (%s)", version.Version, version.GitCommit, version.BuildTime),
	}))

	return echo.WrapHandler(h.Handler())
}
