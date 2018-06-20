package tube

import (
	"fmt"
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
	mb.Registry.MustAddMetricSet("beanstalkd", "tube", New)
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
	cfgwarn.Experimental("The beanstalkd tube metricset is experimental.")

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
	if _, err = conn.Write([]byte("list-tubes\r\n")); err != nil {
		report.Error(err)
		return
	}

	var tubes []string
	var tubesData []string

	if tubesData, err = beanstalkd.ParseResponse(conn); err != nil {
		report.Error(err)
		return
	}

	tubes = beanstalkd.ParseTubes(tubesData)

	if len(tubes) == 0 {
		// nothing to report not tubes available
		return
	}

	for _, tube := range tubes {

		// Fetch the details for the tube
		if _, err = conn.Write([]byte(fmt.Sprintf("stats-tube %s\r\n", tube))); err != nil {
			report.Error(err)
			continue // go to the next one
		}

		var tubeResponse []string
		if tubeResponse, err = beanstalkd.ParseResponse(conn); err != nil {
			report.Error(err)
			continue // skip this one we have problems processing it
		}

		stats := beanstalkd.ParseStats(tubeResponse)
		event, _ := schema.Apply(stats)

		report.Event(mb.Event{
			MetricSetFields: event,
		})

	}

}
