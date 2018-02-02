package tcb

import (
	"log"
	"os"
)

// SilentLogger logger say noting
type SilentLogger struct{}

// Log but say nothing
func (s *SilentLogger) Log(v ...interface{}) {
	return
}

// Logf but say nothing with format either
func (s *SilentLogger) Logf(f string, v ...interface{}) {
	return
}

// DefaultLogger type-default logger
type DefaultLogger struct {
	enabled bool
	Logger  *log.Logger
}

// NewDefaultDisabledLogger Generate New Silent Logger
func NewDefaultDisabledLogger() *DefaultLogger {
	return newDefaultLogger(false)
}

// NewDefaultActiveLogger Generate New Silent Logger
func NewDefaultActiveLogger() *DefaultLogger {
	return newDefaultLogger(true)
}

// newDefaultLogger Generate New Logger
func newDefaultLogger(enabled bool) *DefaultLogger {
	return &DefaultLogger{
		enabled: enabled,
		Logger:  log.New(os.Stderr, "", log.LstdFlags),
	}
}

// LogEnabled return Log-enabled or not
func (d *DefaultLogger) LogEnabled() bool {
	if d == nil {
		return false
	}
	return d.enabled
}

// Log logging with log-package
func (d *DefaultLogger) Log(v ...interface{}) {
	if d.LogEnabled() {
		d.Logger.Println(v...)
	}
}

// Logf logging with format using log-package
func (d *DefaultLogger) Logf(f string, v ...interface{}) {
	if d.LogEnabled() {
		d.Logger.Printf(f, v...)
	}
}
