package info

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
		"current-jobs-urgent", "current-jobs-ready", "current-jobs-reserved", "current-jobs-delayed", "current-jobs-buried",
		"cmd-put", "cmd-peek", "cmd-peek-ready", "cmd-peek-delayed", "cmd-peek-buried", "cmd-reserve",
		"cmd-reserve-with-timeout", "cmd-delete", "cmd-release", "cmd-use", "cmd-watch", "cmd-ignore",
		"cmd-bury", "cmd-kick", "cmd-touch", "cmd-stats", "cmd-stats-job", "cmd-stats-tube", "cmd-list-tubes",
		"cmd-list-tube-used", "cmd-list-tubes-watched", "cmd-pause-tube", "job-timeouts", "total-jobs", "max-job-size",
		"current-tubes", "current-connections", "current-producers", "current-workers", "current-waiting", "total-connections",
		"pid", "version", "rusage-utime", "rusage-stime", "uptime", "binlog-oldest-index", "binlog-current-index",
		"binlog-records-migrated", "binlog-records-written", "binlog-max-size", "id", "hostname"}

	sort.Strings(reqKeys)
	sort.Strings(keys)

	assert.EqualValues(t, reqKeys, keys)
}
