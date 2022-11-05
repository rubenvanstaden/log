package log_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/rubenvanstaden/log"
	"github.com/rubenvanstaden/test"
)

const (
	rgxPID  = `[0-9]+`
	rgxdate = `[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]`
	rgxtime = `[0-9][0-9]:[0-9][0-9]:[0-9][0-9]`
)

type tester struct {
	message string
	pattern string // regexp that log output must match
}

func TestUnitDebug(t *testing.T) {

	tests := []tester{
		{
			message: "hello, world!",
			pattern: fmt.Sprintf("test: pid=%s %s %s DEBUG: hello, world!\n$", rgxPID, rgxdate, rgxtime),
		},
	}

	for _, tc := range tests {

		var buf bytes.Buffer

		logger := log.NewLogger("test", &buf)
		logger.Debug(tc.message)

		act := buf.String()

		matched, err := regexp.MatchString(tc.pattern, act)
		test.Ok(t, err)
		test.Assert(t, matched, fmt.Sprintf("pattern match failed: %q", act))
	}
}
