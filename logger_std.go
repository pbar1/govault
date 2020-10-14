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

func (l *StdLogger) Panic(msg string, items ...interface{}) {
	if l.Level <= LevelPanic {
		items = append(items, "msg", msg)
		log.Println(items)
		panic(msg)
	}
}

func (l *StdLogger) Fatal(msg string, items ...interface{}) {
	if l.Level <= LevelFatal {
		items = append(items, "msg", msg)
		log.Fatalln(items)
	}
}

func (l *StdLogger) Error(msg string, items ...interface{}) {
	if l.Level <= LevelError {
		items = append(items, "msg", msg)
		log.Println(items)
	}
}

func (l *StdLogger) Warn(msg string, items ...interface{}) {
	if l.Level <= LevelWarn {
		items = append(items, "msg", msg)
		log.Println(items)
	}
}

func (l *StdLogger) Info(msg string, items ...interface{}) {
	if l.Level <= LevelInfo {
		items = append(items, "msg", msg)
		log.Println(items)
	}
}

func (l *StdLogger) Debug(msg string, items ...interface{}) {
	if l.Level <= LevelDebug {
		items = append(items, "msg", msg)
		log.Println(items)
	}
}

func (l *StdLogger) Trace(msg string, items ...interface{}) {
	if l.Level <= LevelTrace {
		items = append(items, "msg", msg)
		log.Println(items)
	}
}
