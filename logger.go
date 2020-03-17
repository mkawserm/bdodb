package bdodb

import (
	"log"
	"os"
)

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

var DefaultLogger = &defaultLog{Logger: log.New(os.Stderr, "bdodb ", log.LstdFlags)}
