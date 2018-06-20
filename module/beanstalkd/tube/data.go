package tube

import (
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstrstr"
)

var schema = s.Schema{
	"name":                  c.Str("name"),
	"current-jobs-buried":   c.Int("current-jobs-buried"),
	"current-jobs-delayed":  c.Int("current-jobs-delayed"),
	"current-jobs-ready":    c.Int("current-jobs-ready"),
	"current-jobs-reserved": c.Int("current-jobs-reserved"),
	"current-jobs-urgent":   c.Int("current-jobs-urgent"),
	"total-jobs":            c.Int("total-jobs"),
	"cmd-delete":            c.Int("cmd-delete"),
	"cmd-pause-tube":        c.Int("cmd-pause-tube"),
	"pause":                 c.Int("pause"),
	"pause-time-left":       c.Int("pause-time-left"),
	"current-using":         c.Int("current-using"),
	"current-watching":      c.Int("current-watching"),
	"current-waiting":       c.Int("current-waiting"),
}
