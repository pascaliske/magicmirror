package proxy

import (
	"net/url"
	"strings"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
)

func Handler(server *echo.Echo, target string) echo.MiddlewareFunc {
	logger.Debug("Using %s proxy for %s", config.GetString("Environment"), color.CyanString(target))

	url, _ := url.Parse(target)

	return middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRandomBalancer([]*middleware.ProxyTarget{
			{
				URL: url,
			},
		}),
		Skipper: func(c echo.Context) bool {
			if skip := strings.Contains(c.Request().RequestURI, "health"); skip {
				return skip
			}
			if skip := strings.Contains(c.Request().RequestURI, "socket"); skip {
				return skip
			}
			if skip := strings.Contains(c.Request().RequestURI, config.GetString("Metrics.Path")); skip {
				return skip
			}
			return false
		},
	})
}
