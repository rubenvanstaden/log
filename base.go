package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
)

type Base interface {
	// Debug logs a message at Debug level.
	Debug(args ...interface{})

	// Info logs a message at Info level.
	Info(args ...interface{})

	// Warn logs a message at Warning level.
	Warn(args ...interface{})

	// Error logs a message at Error level.
	Error(args ...interface{})

	// Fatal logs a message at Fatal level
	Fatal(args ...interface{})
}

// Wrapper object around log.Logger from the standard library.
type base struct {
	*stdlog.Logger
}

func newBase(name string, out io.Writer) *base {

	prefix := fmt.Sprintf("%s: pid=%d ", name, os.Getpid())

	return &base{
		stdlog.New(out, prefix, stdlog.Ldate|stdlog.Ltime),
	}
}

func (self *base) Debug(args ...interface{}) {
	self.prefixPrint("DEBUG: ", args...)
}

func (self *base) Info(args ...interface{}) {
	self.prefixPrint("INFO: ", args...)
}

func (self *base) Warn(args ...interface{}) {
	self.prefixPrint("WARN: ", args...)
}

func (self *base) Error(args ...interface{}) {
	self.prefixPrint("ERROR: ", args...)
}

// Process will exit with status set to 1.
func (self *base) Fatal(args ...interface{}) {
	self.prefixPrint("FATAL: ", args...)
	os.Exit(1)
}

func (self *base) prefixPrint(prefix string, args ...interface{}) {
	args = append([]interface{}{prefix}, args...)
	self.Print(args...)
}
