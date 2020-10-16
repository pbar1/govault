package govault

import "os"

type DiscardLogger struct{}

func NewDiscardLogger() *DiscardLogger { return &DiscardLogger{} }

func (l *DiscardLogger) Panic(items ...interface{}) { panic(nil) }
func (l *DiscardLogger) Fatal(items ...interface{}) { os.Exit(1) }
func (l *DiscardLogger) Error(items ...interface{}) {}
func (l *DiscardLogger) Warn(items ...interface{})  {}
func (l *DiscardLogger) Info(items ...interface{})  {}
func (l *DiscardLogger) Debug(items ...interface{}) {}
func (l *DiscardLogger) Trace(items ...interface{}) {}
