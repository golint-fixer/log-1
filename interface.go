package log

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

// Interface represents the API of both Logger and Entry.
type Interface interface {
	WithFields(fields log.Fields) *log.Entry
	WithField(key string, value interface{}) *log.Entry
	WithError(err error) *log.Entry
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Debugf(msg string, v ...interface{})
	Infof(msg string, v ...interface{})
	Warnf(msg string, v ...interface{})
	Errorf(msg string, v ...interface{})
	Fatalf(msg string, v ...interface{})
}
