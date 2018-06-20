package beanstalkd

import (
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingNotFound(t *testing.T) {

	data := "NOT_FOUND\r\n"

	wg := sync.WaitGroup{}

	r, w := io.Pipe()
	// so we can simulate real input
	wg.Add(1)
	go func() {
		defer wg.Done()
		n, err := w.Write([]byte(data))
		assert.NoError(t, err)
		assert.True(t, n > 0)
	}()

	response, err := ParseResponse(r)
	assert.Empty(t, response)
	assert.EqualError(t, err, errObjectNotFound)

	wg.Wait()

}

func TestParsingInvalidStatusCode(t *testing.T) {

	data := "INVALID_STATUS_CODE\r\n"

	wg := sync.WaitGroup{}

	r, w := io.Pipe()
	// so we can simulate real input
	wg.Add(1)
	go func() {
		defer wg.Done()
		n, err := w.Write([]byte(data))
		assert.NoError(t, err)
		assert.True(t, n > 0)
	}()

	response, err := ParseResponse(r)
	assert.Empty(t, response)
	assert.EqualError(t, err, errInvalidResponseCode)

	wg.Wait()

}

func TestParsingTubes(t *testing.T) {

	data := `OK 517
---
- default
- 3bp37KZqDSI4hfxJPZ36L
- oUcUN8h3A2fLeEhaJ
- ukClKcv7QKZCV1q3LaKMLjrwJ
- 7a0bFI3h4v4QJjh8D9Vfn
- 0NnHMzLHjw2IaV9EFvwRCH3gd
- YJR4ILLPq5ho5l
- BALKnl3y2uGA0533dHHbw66l9NoqALKyxI
- XzwUGygudtjaFt2KxAoIxv
- BwIWzo7bmP2kuyL
- A99u4ZpfmKObVtQJ9EOM
- pYWU2oYxKqNrbXldadRuUrU3ydKUxCZqpsuU
- VwsEXaP5qJuORECsHM
- j5cvhCgGe4ghUBB
- aCoykPT1wY6jRY
- XlnIgpG2nagMboYN
- HwHnK0NAbyaodfeZPUD7iC2CncNerog7mQ0uyjxp3i7EpOI3il3Hw2IK
- t2rHLlsHybsZyQchru67
- 83dENJUI3Ny6BXR9Q4y3UpTwG1
- qvQ95D
- dzexeBoCBS1ie0xRqQpyFX

`

	tubes := []string{
		"default",
		"3bp37KZqDSI4hfxJPZ36L",
		"oUcUN8h3A2fLeEhaJ",
		"ukClKcv7QKZCV1q3LaKMLjrwJ",
		"7a0bFI3h4v4QJjh8D9Vfn",
		"0NnHMzLHjw2IaV9EFvwRCH3gd",
		"YJR4ILLPq5ho5l",
		"BALKnl3y2uGA0533dHHbw66l9NoqALKyxI",
		"XzwUGygudtjaFt2KxAoIxv",
		"BwIWzo7bmP2kuyL",
		"A99u4ZpfmKObVtQJ9EOM",
		"pYWU2oYxKqNrbXldadRuUrU3ydKUxCZqpsuU",
		"VwsEXaP5qJuORECsHM",
		"j5cvhCgGe4ghUBB",
		"aCoykPT1wY6jRY",
		"XlnIgpG2nagMboYN",
		"HwHnK0NAbyaodfeZPUD7iC2CncNerog7mQ0uyjxp3i7EpOI3il3Hw2IK",
		"t2rHLlsHybsZyQchru67",
		"83dENJUI3Ny6BXR9Q4y3UpTwG1",
		"qvQ95D",
		"dzexeBoCBS1ie0xRqQpyFX",
	}

	wg := sync.WaitGroup{}

	r, w := io.Pipe()
	// so we can simulate real input
	wg.Add(1)
	go func() {
		defer wg.Done()
		n, err := w.Write([]byte(data))
		assert.NoError(t, err)
		assert.True(t, n > 0)
	}()

	response, err := ParseResponse(r)
	assert.NotEmpty(t, response)
	assert.NoError(t, err)

	assert.EqualValues(t, tubes, ParseTubes(response))

	wg.Wait()

}

func TestParsingStats(t *testing.T) {

	data := `OK 295
---
name: tube-name
current-jobs-urgent: 0
current-jobs-ready: 402152
current-jobs-reserved: 6
current-jobs-delayed: 1
current-jobs-buried: 0
total-jobs: 2437693
current-using: 6
current-watching: 6
current-waiting: 0
cmd-delete: 2035534
cmd-pause-tube: 0
pause: 0
pause-time-left: 0

`
	statsMap := map[string]interface{}{
		"name":                  "tube-name",
		"current-jobs-urgent":   "0",
		"current-jobs-ready":    "402152",
		"current-jobs-reserved": "6",
		"current-jobs-delayed":  "1",
		"current-jobs-buried":   "0",
		"total-jobs":            "2437693",
		"current-using":         "6",
		"current-watching":      "6",
		"current-waiting":       "0",
		"cmd-delete":            "2035534",
		"cmd-pause-tube":        "0",
		"pause":                 "0",
		"pause-time-left":       "0",
	}

	wg := sync.WaitGroup{}

	r, w := io.Pipe()
	// so we can simulate real input
	wg.Add(1)
	go func() {
		defer wg.Done()
		n, err := w.Write([]byte(data))
		assert.NoError(t, err)
		assert.True(t, n > 0)
	}()

	response, err := ParseResponse(r)
	assert.NotEmpty(t, response)
	assert.NoError(t, err)

	parsedStats := ParseStats(response)
	assert.EqualValues(t, statsMap, parsedStats, "Missing items")

	wg.Wait()

}

func TestParseStatsGlobal(t *testing.T) {

	data := `OK 1088
---
current-jobs-urgent: 1
current-jobs-ready: 206760
current-jobs-reserved: 132
current-jobs-delayed: 67
current-jobs-buried: 49
cmd-put: 735048213
cmd-peek: 0
cmd-peek-ready: 34138
cmd-peek-delayed: 197
cmd-peek-buried: 15477206
cmd-reserve: 0
cmd-reserve-with-timeout: 1732757068
cmd-delete: 734982462
cmd-release: 10759164
cmd-use: 296922774
cmd-watch: 12066526
cmd-ignore: 12066274
cmd-bury: 24975966
cmd-kick: 0
cmd-touch: 0
cmd-stats: 279780
cmd-stats-job: 758392831
cmd-stats-tube: 2435595864
cmd-list-tubes: 1511207
cmd-list-tube-used: 0
cmd-list-tubes-watched: 0
cmd-pause-tube: 0
job-timeouts: 10877
total-jobs: 735048213
max-job-size: 5000000
current-tubes: 132
current-connections: 525
current-producers: 282
current-workers: 517
current-waiting: 0
total-connections: 2294000
pid: 16350
version: 1.10
rusage-utime: 117714.392000
rusage-stime: 220129.108000
uptime: 1398953
binlog-oldest-index: 801671
binlog-current-index: 801764
binlog-records-migrated: 226960840
binlog-records-written: 1732728028
binlog-max-size: 10485760
id: 3115ef5b2496e5c8
hostname: abcdefg-beanstalk

`

	statsMap := map[string]interface{}{
		"current-jobs-urgent":      "1",
		"current-jobs-ready":       "206760",
		"current-jobs-reserved":    "132",
		"current-jobs-delayed":     "67",
		"current-jobs-buried":      "49",
		"cmd-put":                  "735048213",
		"cmd-peek":                 "0",
		"cmd-peek-ready":           "34138",
		"cmd-peek-delayed":         "197",
		"cmd-peek-buried":          "15477206",
		"cmd-reserve":              "0",
		"cmd-reserve-with-timeout": "1732757068",
		"cmd-delete":               "734982462",
		"cmd-release":              "10759164",
		"cmd-use":                  "296922774",
		"cmd-watch":                "12066526",
		"cmd-ignore":               "12066274",
		"cmd-bury":                 "24975966",
		"cmd-kick":                 "0",
		"cmd-touch":                "0",
		"cmd-stats":                "279780",
		"cmd-stats-job":            "758392831",
		"cmd-stats-tube":           "2435595864",
		"cmd-list-tubes":           "1511207",
		"cmd-list-tube-used":       "0",
		"cmd-list-tubes-watched":   "0",
		"cmd-pause-tube":           "0",
		"job-timeouts":             "10877",
		"total-jobs":               "735048213",
		"max-job-size":             "5000000",
		"current-tubes":            "132",
		"current-connections":      "525",
		"current-producers":        "282",
		"current-workers":          "517",
		"current-waiting":          "0",
		"total-connections":        "2294000",
		"pid":                      "16350",
		"version":                  "1.10",
		"rusage-utime":             "117714.392000",
		"rusage-stime":             "220129.108000",
		"uptime":                   "1398953",
		"binlog-oldest-index":      "801671",
		"binlog-current-index":     "801764",
		"binlog-records-migrated":  "226960840",
		"binlog-records-written":   "1732728028",
		"binlog-max-size":          "10485760",
		"id":                       "3115ef5b2496e5c8",
		"hostname":                 "abcdefg-beanstalk",
	}

	wg := sync.WaitGroup{}

	r, w := io.Pipe()
	// so we can simulate real input
	wg.Add(1)
	go func() {
		defer wg.Done()
		n, err := w.Write([]byte(data))
		assert.NoError(t, err)
		assert.True(t, n > 0)
	}()

	response, err := ParseResponse(r)
	assert.NotEmpty(t, response)
	assert.NoError(t, err)

	parsedStats := ParseStats(response)
	assert.EqualValues(t, statsMap, parsedStats, "Missing items")

	wg.Wait()

}
