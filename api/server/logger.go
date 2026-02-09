package server

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server Server) setupLogger() {
	// configure logger middleware
	server.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		// enable log values
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogHost:     true,
		LogRemoteIP: true,
		LogError:    true,
		HandleError: true,

		// construct log line based on values
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				slog.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("host", v.Host),
					slog.String("remote_ip", v.RemoteIP),
				)
			} else {
				slog.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("host", v.Host),
					slog.String("remote_ip", v.RemoteIP),
					slog.String("error", v.Error.Error()),
				)
			}

			return nil
		},

		// skip non error logs
		Skipper: func(c echo.Context) bool {
			return c.Response().Status >= 200 && c.Response().Status <= 299
		},
	}))
}
