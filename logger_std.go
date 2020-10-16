package govault

import (
	"log"
)

const (
	LevelPanic = 5
	LevelFatal = 4
	LevelError = 3
	LevelWarn  = 2
	LevelInfo  = 1
	LevelDebug = 0
	LevelTrace = -1
)

type StdLogger struct {
	Level int
}

func NewStdLogger() *StdLogger {
	return &StdLogger{Level: LevelInfo}
}

func (l *StdLogger) Panic(items ...interface{}) {
	if l.Level <= LevelPanic {
		panic(items)
	}
}

func (l *StdLogger) Fatal(items ...interface{}) {
	if l.Level <= LevelFatal {
		log.Fatalln(items...)
	}
}

func (l *StdLogger) Error(items ...interface{}) {
	if l.Level <= LevelError {
		log.Println(items...)
	}
}

func (l *StdLogger) Warn(items ...interface{}) {
	if l.Level <= LevelWarn {
		log.Println(items...)
	}
}

func (l *StdLogger) Info(items ...interface{}) {
	if l.Level <= LevelInfo {
		log.Println(items...)
	}
}

func (l *StdLogger) Debug(items ...interface{}) {
	if l.Level <= LevelDebug {
		log.Println(items...)
	}
}

func (l *StdLogger) Trace(items ...interface{}) {
	if l.Level <= LevelTrace {
		log.Println(items...)
	}
}
