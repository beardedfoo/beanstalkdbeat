package info

import "github.com/ilijamt/beanstalkdbeat/module/beanstalkd"

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "beanstalkd",
		"metricsets": []string{"info"},
		"hosts":      []string{beanstalkd.GetBeanstalkDEnvHost() + ":" + beanstalkd.GetBeanstalkDEnvPort()},
	}
}
