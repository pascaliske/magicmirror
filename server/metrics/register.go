package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// general

var BuildInfo = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Subsystem: "magicmirror",
		Name:      "build_info",
		Help:      "Current version number and build information.",
	},
	[]string{"version", "git_commit", "build_time", "go_version", "platform"},
)

var UptimeSeconds = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Subsystem: "magicmirror",
		Name:      "uptime_seconds",
		Help:      "Seconds since last successfull startup.",
	},
	[]string{},
)

// config reloads

var ConfigReloadsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Subsystem: "magicmirror",
		Name:      "config_reloads_total",
		Help:      "Counter of total successfull config file reloads.",
	},
	[]string{},
)

var ConfigReloadsFailureTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Subsystem: "magicmirror",
		Name:      "config_reloads_failure_total",
		Help:      "Counter of total failed config file reloads.",
	},
	[]string{},
)

var ConfigLastReloadSuccess = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Subsystem: "magicmirror",
		Name:      "config_last_reload_success",
		Help:      "UNIX timestamp of last successfull config file reload.",
	},
	[]string{},
)

var ConfigLastReloadFailure = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Subsystem: "magicmirror",
		Name:      "config_last_reload_failure",
		Help:      "UNIX timestamp of last failed config file reload.",
	},
	[]string{},
)

// http

var RequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Subsystem: "magicmirror",
		Name:      "requests_total",
		Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
	},
	[]string{"code", "method", "url", "host"},
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

// socket

var SocketConnections = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Subsystem: "magicmirror",
		Name:      "socket_connections",
		Help:      "Number of socket clients currently connected.",
	},
	[]string{},
)
