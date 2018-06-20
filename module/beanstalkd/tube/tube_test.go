package tube

import (
	"testing"

	"sort"

	"github.com/stretchr/testify/assert"

	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

func TestFetchEventContents(t *testing.T) {

	f := mbtest.NewReportingMetricSetV2(t, getConfig())

	events, errs := mbtest.ReportingFetchV2(f)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
	if len(events) == 0 {
		t.Fatal("no events received")
	}

	for _, e := range events {
		if e.Error != nil {
			t.Fatalf("received error: %+v", e.Error)
		}
	}
	if len(events) == 0 {
		t.Fatal("received no events")
	}

	e := events[0]

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), e)

	keys := make([]string, 0, len(e.MetricSetFields))
	for k := range e.MetricSetFields {
		keys = append(keys, k)
	}

	reqKeys := []string{
		"name", "current-jobs-buried", "current-jobs-delayed",
		"current-jobs-ready", "current-jobs-reserved", "current-jobs-urgent",
		"total-jobs", "cmd-delete", "cmd-pause-tube",
		"pause", "pause-time-left", "current-using",
		"current-watching", "current-waiting"}

	sort.Strings(reqKeys)
	sort.Strings(keys)

	assert.EqualValues(t, reqKeys, keys)

}
