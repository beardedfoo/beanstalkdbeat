package beanstalkd

import (
	"os"
)

// Helper functions for testing used in the redis metricsets

// GetBeanstalkDEnvHost returns the hostname of the BeanstalkD server to use for testing.
// It reads the value from the BEANSTALKD_HOST environment variable and returns
// 127.0.0.1 if it is not set.
func GetBeanstalkDEnvHost() string {
	host := os.Getenv("BEANSTALKD_HOST")

	if len(host) == 0 {
		host = "127.0.0.1"
	}
	return host
}

// GetBeanstalkDEnvPort returns the port of the BeanstalkD server to use for testing.
// It reads the value from the BEANSTALKD_PORT environment variable and returns
// 11300 if it is not set.
func GetBeanstalkDEnvPort() string {
	port := os.Getenv("BEANSTALKD_PORT")

	if len(port) == 0 {
		port = "11300"
	}
	return port
}
