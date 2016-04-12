// Package log implements a featured and structured generic logger
// interface designed to be used in core and third-party vinxi components.
package log

import (
	"io"
	"sync"
	"time"

	log "gopkg.in/Sirupsen/logrus.v0"
)

// mu provides thread synchronization for the public singleton API.
var mu = sync.Mutex{}

// Logger represents the standard default built-in logger (equivalent to default Logrus logger).
var Logger = log.New()

// New creates a new Logrus based logger.
func New() *log.Logger {
	return log.New()
}

// WithFields returns a new entry with `fields` set.
func WithFields(fields log.Fields) *log.Entry {
	return Logger.WithFields(fields)
}

// WithField returns a new entry with the `key` and `value` set.
func WithField(key string, value interface{}) *log.Entry {
	return Logger.WithField(key, value)
}

// WithError returns a new entry with the "error" set to `err`.
func WithError(err error) *log.Entry {
	return Logger.WithError(err)
}

// Debug level message.
func Debug(msg string) {
	Logger.Debug(msg)
}

// Info level message.
func Info(msg string) {
	Logger.Info(msg)
}

// Warn level message.
func Warn(msg string) {
	Logger.Warn(msg)
}

// Error level message.
func Error(msg string) {
	Logger.Error(msg)
}

// Fatal level message, followed by an exit.
func Fatal(msg string) {
	Logger.Fatal(msg)
}

// Debugf level formatted message.
func Debugf(msg string, v ...interface{}) {
	Logger.Debugf(msg, v...)
}

// Infof level formatted message.
func Infof(msg string, v ...interface{}) {
	Logger.Infof(msg, v...)
}

// Warnf level formatted message.
func Warnf(msg string, v ...interface{}) {
	Logger.Warnf(msg, v...)
}

// Errorf level formatted message.
func Errorf(msg string, v ...interface{}) {
	Logger.Errorf(msg, v...)
}

// Fatalf level formatted message, followed by an exit.
func Fatalf(msg string, v ...interface{}) {
	Logger.Fatalf(msg, v...)
}

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func Trace(msg string) *log.Entry {
	e := log.NewEntry(Logger)
	e.Info(msg)
	v := e.WithFields(e.Data)
	v.Message = msg
	v.Time = time.Now()
	return v
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	Logger.Out = out
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter log.Formatter) {
	mu.Lock()
	defer mu.Unlock()
	Logger.Formatter = formatter
}

// SetLevel sets the standard logger level.
func SetLevel(level log.Level) {
	mu.Lock()
	defer mu.Unlock()
	Logger.Level = level
}

// GetLevel returns the standard logger level.
func GetLevel() log.Level {
	mu.Lock()
	defer mu.Unlock()
	return Logger.Level
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook log.Hook) {
	mu.Lock()
	defer mu.Unlock()
	Logger.Hooks.Add(hook)
}
