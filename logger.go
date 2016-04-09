package framework

import (
	"io"
)

type Entry interface {
	LogWriter
}

type EntryFormatter interface {
	Format(string, ...interface{}) string
}

type LogWriter interface {
	Print(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Warning(string, ...interface{})
	Fatal(string, ...interface{})
	With(...interface{}) Entry
}

type Logger interface {
	LogWriter
	Formatter(EntryFormatter)
	DefaultWith(...interface{}) LogWriter
	Writers(...io.Writer)
	Logger(Context) Logger
}
