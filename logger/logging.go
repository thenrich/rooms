package logger

import (
	aelog "google.golang.org/appengine/log"

	"golang.org/x/net/context"
)

var DefaultLogger = &Logger{}

// Logger provides indirection for AppEngine
type Logger struct {
	ctx context.Context
}

// Set the Context for the logs
func (l *Logger) SetContext(c context.Context) {
	l.ctx = c
}

// Calls AppEngine's Infof
func (l *Logger) Infof(format string, args ...interface{}) {
	aelog.Infof(l.ctx, format, args)
}

// Calls AppEngine's Errorf
func (l *Logger) Errorf(format string, args ...interface{}) {
	aelog.Errorf(l.ctx, format, args)
}