package ownmetrics

import (
	"log"
	"math"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
)

const namespace = "go_customs_stats"

type CustomerCollector struct {
	cpuUsage  *prometheus.Desc
	uptime    *prometheus.Desc
	startTime time.Time
}

func (e *CustomerCollector) cpuStatus() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}

	return math.Ceil(percent[0])
}

func (e *CustomerCollector) upTime() float64 {
	return time.Since(e.startTime).Seconds()
}

func NewCustomerCollector(startTime time.Time) *CustomerCollector {
	return &CustomerCollector{
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

func (e *CustomerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.cpuUsage
	ch <- e.uptime
}

func (e *CustomerCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		e.cpuUsage,
		prometheus.GaugeValue,
		e.cpuStatus())
	ch <- prometheus.MustNewConstMetric(
		e.uptime,
		prometheus.CounterValue,
		e.upTime())
}
