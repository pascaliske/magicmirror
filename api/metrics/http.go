package metrics

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type HttpMetrics struct {
	Context echo.Context
	Start   time.Time
}

func NewHttpMetrics(c echo.Context) HttpMetrics {
	return HttpMetrics{
		Context: c,
		Start:   time.Now(),
	}
}

func (metrics HttpMetrics) Status(err error) string {
	// initial response status
	status := metrics.Context.Response().Status

	// determine response status from middleware chain
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

func (metrics HttpMetrics) Duration() float64 {
	return float64(time.Since(metrics.Start)) / float64(time.Second)
}

func (metrics HttpMetrics) RequestSize() float64 {
	size := 0

	// add url length
	if metrics.Context.Request().URL != nil {
		size = len(metrics.Context.Request().URL.Path)
	}

	// add method, protocol and host lengths
	size += len(metrics.Context.Request().Method)
	size += len(metrics.Context.Request().Proto)
	size += len(metrics.Context.Request().Host)

	// add header lengths
	for name, values := range metrics.Context.Request().Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}

	// add content length value
	if metrics.Context.Request().ContentLength != -1 {
		size += int(metrics.Context.Request().ContentLength)
	}

	return float64(size)
}

func (metrics HttpMetrics) ResponseSize() float64 {
	return float64(metrics.Context.Response().Size)
}

func (metrics HttpMetrics) Labels(err error) []string {
	return []string{
		metrics.Status(err),              // code
		metrics.Context.Request().Method, // method
		metrics.Context.Path(),           // url
	}
}

func (metrics HttpMetrics) LabelsWithHost(err error) []string {
	return append(metrics.Labels(err), metrics.Context.Request().Host)
}
