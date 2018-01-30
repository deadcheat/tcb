package tcb

import "log"

// SilentLogger logger say noting
type SilentLogger struct{}

// Log but say nothing
func (s *SilentLogger) Log(v ...interface{}) {
	if s == nil {
		return
	}
}

// Logf but say nothing with format either
func (s *SilentLogger) Logf(f string, v ...interface{}) {
	if s == nil {
		return
	}
}

// DefaultLogger type-default logger
type DefaultLogger struct {
	enabled bool
}

// NewDefaultDisabledLogger Generate New Silent Logger
func NewDefaultDisabledLogger() Loggerable {
	return newDefaultLogger(false)
}

// NewDefaultActiveLogger Generate New Silent Logger
func NewDefaultActiveLogger() Loggerable {
	return newDefaultLogger(false)
}

// newDefaultLogger Generate New Logger
func newDefaultLogger(enabled bool) Loggerable {
	return &DefaultLogger{
		enabled: enabled,
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
		log.Println(v...)
	}
}

// Logf logging with format using log-package
func (d *DefaultLogger) Logf(f string, v ...interface{}) {
	if d.LogEnabled() {
		log.Printf(f, v...)
	}
}
