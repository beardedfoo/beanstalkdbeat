package info

import (
	"net"

	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/ilijamt/beanstalkdbeat/module/beanstalkd"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("beanstalkd", "info", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The beanstalkd info metricset is experimental.")

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {

	var conn net.Conn
	var err error

	if conn, err = net.DialTimeout("tcp", m.Host(), m.Module().Config().Timeout); err != nil {
		report.Error(err)
		return
	}
	defer conn.Close()

	// Fetch a list of all the tubes we have
	if _, err = conn.Write([]byte("stats\r\n")); err != nil {
		report.Error(err)
		return
	}

	var stats map[string]interface{}
	var statsData []string

	if statsData, err = beanstalkd.ParseResponse(conn); err != nil {
		report.Error(err)
		return
	}

	stats = beanstalkd.ParseStats(statsData)
	event, _ := schema.Apply(stats)

	report.Event(mb.Event{
		MetricSetFields: event,
	})

}
