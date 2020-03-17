package bdodb

import (
	"log"
	"os"
)

// Logger interface
type Logger interface {
	Errorf(f string, v ...interface{})
	Warningf(f string, v ...interface{})
	Infof(f string, v ...interface{})
	Debugf(f string, v ...interface{})
}

type defaultLog struct {
	*log.Logger
}

func (l *defaultLog) Errorf(f string, v ...interface{}) {
	l.Printf("ERROR: "+f, v...)
}

func (l *defaultLog) Warningf(f string, v ...interface{}) {
	l.Printf("WARNING: "+f, v...)
}

func (l *defaultLog) Infof(f string, v ...interface{}) {
	l.Printf("INFO: "+f, v...)
}

func (l *defaultLog) Debugf(f string, v ...interface{}) {
	l.Printf("DEBUG: "+f, v...)
}

// DefaultLogger is the default logger for bdodb
// Set different logger to modify the logging behavior
var DefaultLogger Logger = &defaultLog{Logger: log.New(os.Stderr, "bdodb ", log.LstdFlags)}
