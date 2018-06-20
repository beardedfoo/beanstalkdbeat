// +build integration

package info

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/tests/compose"
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestFetch(t *testing.T) {
	compose.EnsureUp(t, "beanstalkd")

	f := mbtest.NewReportingMetricSetV2(t, getConfig())

	events, errs := mbtest.ReportingFetchV2(f)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
	if len(events) == 0 {
		t.Fatal("no events received")
	}

	assert.NotNil(t, events)
	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), events)
}
