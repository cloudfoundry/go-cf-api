package custommetrics_test

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/metrics/custommetrics"
)

func TestCustomCollector(t *testing.T) {
	reg := prometheus.NewRegistry()
	{
		if err := reg.Register(custommetrics.NewCustomCollector(time.Now().UTC())); err != nil {
			t.Fatal(err)
		}
	}

	mfs, err := reg.Gather()
	if err != nil {
		t.Fatal(err)
	}

	names := []string{
		"go_custom_stats_cpu_usage_total",
		"go_custom_stats_uptime_seconds_total",
	}

	type result struct {
		found bool
	}
	results := make(map[string]result)
	for _, name := range names {
		results[name] = result{found: false}
	}
	for _, mf := range mfs {
		for _, name := range names {
			if name == mf.GetName() {
				results[name] = result{found: true}
				break
			}
		}
	}

	for name, result := range results {
		if !result.found {
			t.Errorf("%s not found", name)
		}
	}
}
