package log

import (
	"fmt"
	"io"
	"sync"
)

func NewLogger(name string, out io.Writer) *Logger {

	base := newBase(name, out)

	return &Logger{
		base:  base,
		level: DebugLevel,
	}
}

// Logger logs message to io.Writer at various log levels.
type Logger struct {
	base  Base
	mu    sync.Mutex
	level Level
}

// SetLevel sets the logger level.
// It panics if v is less than DebugLevel or greater than FatalLevel.
func (self *Logger) SetLevel(v Level) {

	self.mu.Lock()
	defer self.mu.Unlock()

	if v < DebugLevel || v > FatalLevel {
		panic("log: invalid log level")
	}

	self.level = v
}

// checkLevel reports whether logger can log at level v.
func (self *Logger) checkLevel(v Level) bool {

	self.mu.Lock()
	defer self.mu.Unlock()

	return v >= self.level
}

func (self *Logger) Debug(args ...interface{}) {
	if !self.checkLevel(DebugLevel) {
		return
	}
	self.base.Debug(args...)
}

func (self *Logger) Info(args ...interface{}) {
	if !self.checkLevel(InfoLevel) {
		return
	}
	self.base.Info(args...)
}

func (self *Logger) Warn(args ...interface{}) {
	if !self.checkLevel(WarnLevel) {
		return
	}
	self.base.Warn(args...)
}

func (self *Logger) Error(args ...interface{}) {
	if !self.checkLevel(ErrorLevel) {
		return
	}
	self.base.Error(args...)
}

func (self *Logger) Fatal(args ...interface{}) {
	if !self.checkLevel(FatalLevel) {
		return
	}
	self.base.Fatal(args...)
}

func (self *Logger) Debugf(format string, args ...interface{}) {
	self.Debug(fmt.Sprintf(format, args...))
}

func (self *Logger) Infof(format string, args ...interface{}) {
	self.Info(fmt.Sprintf(format, args...))
}

func (self *Logger) Warnf(format string, args ...interface{}) {
	self.Warn(fmt.Sprintf(format, args...))
}

func (self *Logger) Errorf(format string, args ...interface{}) {
	self.Error(fmt.Sprintf(format, args...))
}

func (self *Logger) Fatalf(format string, args ...interface{}) {
	self.Fatal(fmt.Sprintf(format, args...))
}
