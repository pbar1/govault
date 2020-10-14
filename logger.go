package govault

type Logger interface {
	Panic(msg string, items ...interface{})
	Fatal(msg string, items ...interface{})
	Error(msg string, items ...interface{})
	Warn(msg string, items ...interface{})
	Info(msg string, items ...interface{})
	Debug(msg string, items ...interface{})
	Trace(msg string, items ...interface{})
}
