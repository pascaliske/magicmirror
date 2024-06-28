package metrics

import (
	"time"
)

type UptimeMetrics struct {
	startup time.Time
}

func NewUptimeMetrics() UptimeMetrics {
	return UptimeMetrics{startup: time.Now()}
}

func (metrics UptimeMetrics) GetStartup() float64 {
	return float64(metrics.startup.Unix())
}

func (metrics UptimeMetrics) GetUptime() float64 {
	return float64(time.Now().Unix() - metrics.startup.Unix())
}
