package framework

import (
	"net/http"
	"time"
)

type Server interface {
	SetHandler(http.Handler)
	SetLogger(LogWriter)
	SetAddr(string)
	ListenAndServe() error
	ListenAndServeTLS(string, string) error
}

type Valuer interface {
	Values() Values
	SetValue(interface{}, interface{})
	GetValue(interface{}) Value
	Value(interface{}) interface{}
}

type Context interface {
	LogWriter
	Request
	Response
	Valuer
	Operation() Operation
}

type ParamValuer interface {
	Param(string) ParamValue
	ParamOk(string) (ParamValue, bool)
	Params(...string) ParamValues
}

type Request interface {
	Sessioner
	ParamValuer
	Path() string
	Method() string
}

type Operation interface {
	BuildPather
	Route() Route
	GetName() string
	GetMethods() []string
	GetPath() string
	GetFullPath() string
}

type Response interface {
	http.ResponseWriter
	IsIndented() bool
}
type Param interface {
}
type Params []Param
type ParamValue interface {
	RawString() string
}
type ParamValues map[string]ParamValue
type Route interface {
	BuildPather
	Path() string
	FullPath() string
	SetPath(string)
}

