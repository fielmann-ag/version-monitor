package logging

import (
	"testing"
)

var _ Logger = &TestLogger{}

// TestLogger implements a test wrapper for testing.T implementing Logger interface
type TestLogger struct {
	t *testing.T
}

// NewTestLogger returns a new TestLogger
func NewTestLogger(t *testing.T) *TestLogger {
	return &TestLogger{
		t: t,
	}
}

// Debugf implements Logger interface
func (l *TestLogger) Debugf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Infof implements Logger interface
func (l *TestLogger) Infof(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Printf implements Logger interface
func (l *TestLogger) Printf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Warnf implements Logger interface
func (l *TestLogger) Warnf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Warningf implements Logger interface
func (l *TestLogger) Warningf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Errorf implements Logger interface
func (l *TestLogger) Errorf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Fatalf implements Logger interface
func (l *TestLogger) Fatalf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Panicf implements Logger interface
func (l *TestLogger) Panicf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Debug implements Logger interface
func (l *TestLogger) Debug(args ...interface{}) {
	l.t.Log(args...)
}

// Info implements Logger interface
func (l *TestLogger) Info(args ...interface{}) {
	l.t.Log(args...)
}

// Print implements Logger interface
func (l *TestLogger) Print(args ...interface{}) {
	l.t.Log(args...)
}

// Warn implements Logger interface
func (l *TestLogger) Warn(args ...interface{}) {
	l.t.Log(args...)
}

// Warning implements Logger interface
func (l *TestLogger) Warning(args ...interface{}) {
	l.t.Log(args...)
}

// Error implements Logger interface
func (l *TestLogger) Error(args ...interface{}) {
	l.t.Log(args...)
}

// Fatal implements Logger interface
func (l *TestLogger) Fatal(args ...interface{}) {
	l.t.Log(args...)
}

// Panic implements Logger interface
func (l *TestLogger) Panic(args ...interface{}) {
	l.t.Log(args...)
}
