package log_test

import (
	"os"
	"testing"

	"github.com/rubenvanstaden/log"
)

func TestUnitInfo(t *testing.T) {

    logger := log.NewLogger("taskq", os.Stderr)

    logger.Info("Hello world")

}
