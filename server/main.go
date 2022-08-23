package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/health"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/pascaliske/magicmirror/metrics"
	"github.com/pascaliske/magicmirror/proxy"
	"github.com/pascaliske/magicmirror/socket"
)

// build information
var Version string
var GitCommit string
var BuildTime string

// flags
var configPath *string
var checkMode *bool

func init() {
	configPath = flag.String("config", "", "Path to configuration file")
	checkMode = flag.Bool("check", false, "Enable configuration check only mode")
	flag.Parse()
}

func main() {
	// build information
	figure.NewFigure("MagicMirror", "graffiti", true).Print()
	logger.Raw("\nVersion %s @ %s (%s)\n", color.CyanString(Version), color.CyanString(GitCommit), color.CyanString(BuildTime))

	// parse and validate config
	config.ParseAndValidate(*configPath, *checkMode)

	// configure log level
	logger.SetLevel(config.GetString("Log.Level"))
	config.OnChangeSuccess("log-level", func() {
		logger.SetLevel(config.GetString("Log.Level"))
	})

	// setup server
	server := echo.New()
	server.HidePort = true
	server.HideBanner = true
	server.Use(middleware.CORS())
	server.Use(middleware.Secure())
	server.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, ip=${remote_ip}, latency=${latency_human}\n",
		Skipper: func(c echo.Context) bool {
			return c.Response().Status >= 200 && c.Response().Status <= 299
		},
	}))

	// endpoints
	server.GET("/health", health.Handler(server))
	server.GET("/socket", socket.Handler(server))
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		Index: "index.html",
		HTML5: true,
	}))

	// metrics
	if config.GetBool("Metrics.Enabled") {
		server.Use(metrics.Middleware(Version, GitCommit, BuildTime))
		server.GET(config.GetString("Metrics.Path"), metrics.Handler(server))

		// update config reload metric
		config.OnChange("config-reload", func(success bool) {
			if success {
				metrics.ConfigReloadsTotal.WithLabelValues().Inc()
				metrics.ConfigLastReloadSuccess.WithLabelValues().SetToCurrentTime()
			} else {
				metrics.ConfigReloadsFailureTotal.WithLabelValues().Inc()
				metrics.ConfigLastReloadFailure.WithLabelValues().SetToCurrentTime()
			}
		})
	}

	// development proxy
	if config.GetString("Environment") != "production" {
		server.Use(proxy.Handler(server, "http://localhost:4200"))
	}

	// start server
	go listen(server)

	// graceful shutdown
	shutdown(server)
}

func listen(server *echo.Echo) {
	logger.Info("Server is listening on %s", color.CyanString(fmt.Sprintf(":%d", config.GetInt("Port"))))

	// start server
	if err := server.Start(fmt.Sprintf(":%d", config.GetInt("Port"))); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}

func shutdown(server *echo.Echo) {
	// wait for interrupt or terminate signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// clean-up
	defer logger.Raw("\nGracefully shutting down server...")
	defer signal.Stop(quit)
	defer cancel()

	// shutdown server
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}
