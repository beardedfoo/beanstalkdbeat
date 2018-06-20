package beanstalkd

import (
	"bufio"
	"io"
	"strings"

	"github.com/pkg/errors"
)

const (
	responseOK       = "OK"
	responseNotFound = "NOT_FOUND"

	errCannotFetchFirstLine = "couldn't fetch the first line that contains status code"
	errObjectNotFound       = "object not found"
	errInvalidResponseCode  = "invalid response code"
)

func ParseResponse(reader io.Reader) (data []string, err error) {

	scanner := bufio.NewScanner(reader)

	// get the first line of the text as it contains the status code
	if !scanner.Scan() {
		return nil, errors.New(errCannotFetchFirstLine)
	}

	// A generic response for any requests for is in the format of
	// NOT_FOUND\r\n if the tube/job does not exist.
	// OK <bytes>\r\n<data>\r\n

	text := scanner.Text()
	entries := strings.Split(text, " ")
	switch entries[0] {
	case responseOK:
	case responseNotFound:
		return nil, errors.New(errObjectNotFound)
	default:
		// what happened there should be no other option
		return nil, errors.New(errInvalidResponseCode)
	}

	for scanner.Scan() {

		text := scanner.Text()
		if text == "" {
			// we have reached the end of the line
			break
		} else if text == "---" {
			// we don't need to process this line
			continue
		}

		data = append(data, text)
	}

	return data, nil
}

func ParseTubes(data []string) (tubes []string) {

	for _, item := range data {
		if entries := strings.Split(item, " "); len(entries) == 2 {
			tube := strings.TrimSpace(entries[1])
			tubes = append(tubes, tube)
		}
	}

	return tubes
}

func ParseStats(data []string) map[string]interface{} {

	stats := map[string]interface{}{}

	for _, item := range data {
		if entries := strings.Split(item, ":"); len(entries) == 2 {
			key := strings.TrimSpace(entries[0])
			value := strings.TrimSpace(entries[1])
			stats[key] = strings.TrimSpace(value)
		}
	}

	return stats

}
