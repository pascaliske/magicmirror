package proxy

import (
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pascaliske/magicmirror/config"
)

func Handler(cfg config.Config, server *echo.Echo, target string) echo.MiddlewareFunc {
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
			if skip := strings.Contains(c.Request().RequestURI, cfg.Metrics.Path); skip {
				return skip
			}
			return false
		},
	})
}
