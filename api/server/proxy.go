package server

import (
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
)

func (server Server) setupProxy(target string) {
	// skip proxy in production
	if config.GetString("Environment") == "production" {
		logger.Debug("Skipping proxy for %s environment", config.GetString("Environment"))
		return
	}

	logger.Debug("Using proxy for %s", target)

	// parse target url
	url, _ := url.Parse(target)

	// configure proxy middleware
	server.router.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
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
	}))
}
