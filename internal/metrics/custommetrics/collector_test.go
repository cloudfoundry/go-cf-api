//go:build unit
// +build unit

package custommetrics_test

import (
	"testing"
	"time"

	"github.com/cloudfoundry/go-cf-api/internal/metrics/custommetrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestCustomCollector(t *testing.T) {
	reg := prometheus.NewRegistry()
	err := reg.Register(custommetrics.NewCustomCollector(time.Now().UTC()))
	assert.NoError(t, err)

	mfs, err := reg.Gather()
	assert.NoError(t, err)

	metricNames := []string{
		"go_custom_stats_cpu_usage_total",
		"go_custom_stats_uptime_seconds_total",
	}

	type result struct {
		found bool
	}
	MetricSearchResults := make(map[string]result)

	for _, name := range metricNames {
		MetricSearchResults[name] = result{found: false}
	}
	for _, mf := range mfs {
		for _, name := range metricNames {
			if name == mf.GetName() {
				MetricSearchResults[name] = result{found: true}
				break
			}
		}
	}

	for metricName, result := range MetricSearchResults {
		assert.Truef(t, result.found, "This should be true, %s metric is not found", metricName)
	}
}
