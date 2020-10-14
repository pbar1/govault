package govault

import "os"

type DiscardLogger struct{}

func NewDiscardLogger() *DiscardLogger { return &DiscardLogger{} }

func (l *DiscardLogger) Panic(msg string, items ...interface{}) { panic(nil) }
func (l *DiscardLogger) Fatal(msg string, items ...interface{}) { os.Exit(1) }
func (l *DiscardLogger) Error(msg string, items ...interface{}) {}
func (l *DiscardLogger) Warn(msg string, items ...interface{})  {}
func (l *DiscardLogger) Info(msg string, items ...interface{})  {}
func (l *DiscardLogger) Debug(msg string, items ...interface{}) {}
func (l *DiscardLogger) Trace(msg string, items ...interface{}) {}
