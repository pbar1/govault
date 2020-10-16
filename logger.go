package govault

type Logger interface {
	Panic(items ...interface{})
	Fatal(items ...interface{})
	Error(items ...interface{})
	Warn(items ...interface{})
	Info(items ...interface{})
	Debug(items ...interface{})
	Trace(items ...interface{})
}
