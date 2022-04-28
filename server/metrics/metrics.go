package metrics

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/pascaliske/magicmirror/config"
	"github.com/pascaliske/magicmirror/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var startup int64 = time.Now().Unix()

func Setup(Version string, GitCommit string, BuildTime string) {
	GoVersion := runtime.Version()
	Platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	logger.Debug("Metrics endpoint enabled at %s", color.CyanString(config.GetString("Metrics.Path")))

	// set defaults
	BuildInfo.WithLabelValues(Version, GitCommit, BuildTime, GoVersion, Platform).Set(1)
	UptimeSeconds.WithLabelValues().Set(0)
	ConfigReloadsTotal.WithLabelValues().Add(1)
	ConfigReloadsFailureTotal.WithLabelValues().Add(0)
	ConfigLastReloadSuccess.WithLabelValues().Set(float64(startup))
	ConfigLastReloadFailure.WithLabelValues().Set(0)
	SocketClients.WithLabelValues().Set(0)
}

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// skip on metrics path
			if c.Path() == config.GetString("Metrics.Path") {
				return next(c)
			}

			// continue middleware chain
			start := time.Now()
			err := next(c)

			// preprare values
			status := getStatus(c, err)
			method := c.Request().Method
			host := c.Request().Host
			url := c.Path()

			// update metrics
			UptimeSeconds.WithLabelValues().Set(float64(time.Now().Unix() - startup))
			RequestsTotal.WithLabelValues(status, method, host, url).Inc()
			RequestDuration.WithLabelValues(status, method, url).Observe(float64(time.Since(start)) / float64(time.Second))
			RequestSize.WithLabelValues(status, method, url).Observe(float64(computeApproximateRequestSize(c.Request())))
			ResponseSize.WithLabelValues(status, method, url).Observe(float64(c.Response().Size))

			return err
		}
	}
}

func Handler(server *echo.Echo) echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}

func getStatus(c echo.Context, err error) string {
	status := c.Response().Status

	if err != nil {
		var httpError *echo.HTTPError
		if errors.As(err, &httpError) {
			status = httpError.Code
		}
		if status == 0 || status == http.StatusOK {
			status = http.StatusInternalServerError
		}
	}

	return strconv.Itoa(status)
}

func computeApproximateRequestSize(request *http.Request) int {
	size := 0

	// add url length
	if request.URL != nil {
		size = len(request.URL.Path)
	}

	// add method, protocol and host lengths
	size += len(request.Method)
	size += len(request.Proto)
	size += len(request.Host)

	// add header lengths
	for name, values := range request.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}

	// add content length value
	if request.ContentLength != -1 {
		size += int(request.ContentLength)
	}

	return size
}
