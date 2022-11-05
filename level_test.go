package log_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/rubenvanstaden/log"
	"github.com/rubenvanstaden/test"
)

type levelTester struct {
	level log.Level
	op    string
}

func TestUnit_LowerLevels(t *testing.T) {

	// Logger should not log messages at a level lower than the specified level.
	tests := []levelTester{
		// with level one above
		{log.InfoLevel, "Debug"},
		{log.InfoLevel, "Debugf"},
		{log.WarnLevel, "Info"},
		{log.WarnLevel, "Infof"},
		{log.ErrorLevel, "Warn"},
		{log.ErrorLevel, "Warnf"},
		{log.FatalLevel, "Error"},
		{log.FatalLevel, "Errorf"},
		// with skip level
		{log.WarnLevel, "Debug"},
		{log.ErrorLevel, "Infof"},
	}

	for _, tc := range tests {

		var buf bytes.Buffer

		logger := log.NewLogger("test", &buf)
		logger.SetLevel(tc.level)

		switch tc.op {
		case "Debug":
			logger.Debug("hello")
		case "Debugf":
			logger.Debugf("hello, %s", "world")
		case "Info":
			logger.Info("hello")
		case "Infof":
			logger.Infof("hello, %s", "world")
		case "Warn":
			logger.Warn("hello")
		case "Warnf":
			logger.Warnf("hello, %s", "world")
		case "Error":
			logger.Error("hello")
		case "Errorf":
			logger.Errorf("hello, %s", "world")
		default:
			t.Fatalf("unexpected op: %q", tc.op)
		}

		logged := buf.String() == ""
		failMsg := fmt.Sprintf("logger.%s outputted log message when level is set to %v", tc.op, tc.level)
		test.Assert(t, logged, failMsg)
	}
}
