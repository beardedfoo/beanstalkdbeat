/*
Package include imports all Module and MetricSet packages so that they register
their factories with the global registry. This package can be imported in the
main package to automatically register all of the standard supported Metricbeat
modules.
*/
package include

import (
	// This list is automatically generated by `make imports`
	_ "github.com/ilijamt/beanstalkdbeat/module/beanstalkd"
	_ "github.com/ilijamt/beanstalkdbeat/module/beanstalkd/info"
	_ "github.com/ilijamt/beanstalkdbeat/module/beanstalkd/tube"
)
