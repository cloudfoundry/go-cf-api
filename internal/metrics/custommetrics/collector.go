package custommetrics

import (
	"math"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
)

const namespace = "go_custom_stats"

type CustomCollector struct {
	cpuUsage  *prometheus.Desc
	uptime    *prometheus.Desc
	startTime time.Time
}

func (e *CustomCollector) cpuStatus() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0.0, err
	}

	return math.Ceil(percent[0]), nil
}

func (e *CustomCollector) upTime() float64 {
	return time.Since(e.startTime).Seconds()
}

func NewCustomCollector(startTime time.Time) *CustomCollector {
	return &CustomCollector{
		cpuUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "cpu_usage_total"),
			"Current cpu usage in percentage (*)",
			nil,
			nil),
		uptime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "uptime_seconds_total"),
			"Current uptime in seconds (*)",
			nil,
			nil),
		startTime: startTime,
	}
}

func (e *CustomCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.cpuUsage
	ch <- e.uptime
}

func (e *CustomCollector) Collect(ch chan<- prometheus.Metric) {
	percent, err := e.cpuStatus()
	if err != nil {
		log.Error(err)
	}
	ch <- prometheus.MustNewConstMetric(
		e.cpuUsage,
		prometheus.GaugeValue,
		percent)
	ch <- prometheus.MustNewConstMetric(
		e.uptime,
		prometheus.CounterValue,
		e.upTime())
}
