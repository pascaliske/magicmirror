package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var RequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Subsystem: "magicmirror",
		Name:      "requests_total",
		Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
	},
	[]string{"code", "method", "host", "url"},
)

var RequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Subsystem: "magicmirror",
		Name:      "request_duration_seconds",
		Help:      "The HTTP request latencies in seconds.",
	},
	[]string{"code", "method", "url"},
)

var RequestSize = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Subsystem: "magicmirror",
		Name:      "request_size_bytes",
		Help:      "The HTTP request sizes in bytes.",
	},
	[]string{"code", "method", "url"},
)

var ResponseSize = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Subsystem: "magicmirror",
		Name:      "response_size_bytes",
		Help:      "The HTTP response sizes in bytes.",
	},
	[]string{"code", "method", "url"},
)
