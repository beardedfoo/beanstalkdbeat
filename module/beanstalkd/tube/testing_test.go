package tube

import "github.com/ilijamt/beanstalkdbeat/module/beanstalkd"

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "beanstalkd",
		"metricsets": []string{"tube"},
		"hosts":      []string{beanstalkd.GetBeanstalkDEnvHost() + ":" + beanstalkd.GetBeanstalkDEnvPort()},
	}
}
