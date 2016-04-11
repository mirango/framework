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

type Error interface {
	error
	Code() int
	Type() string
	Name() string
}

type Config interface {
	Get(string) Value
}

type Valuer interface {
	Values() Values
	SetValue(interface{}, interface{})
	GetValue(interface{}) Value
	Value(interface{}) interface{}
}

type ParamValuer interface {
	Param(string) ParamValue
	ParamOk(string) (ParamValue, bool)
	Params(...string) ParamValues
}

type Context interface {
	LogWriter
	Request
	Response
	Valuer
	Operation() Operation
	Locale() Locale
	Id() int64
	User() User
}

type Request interface {
	Sessioner
	ParamValuer
	Path() string
	Method() string
}

type Response interface {
	http.ResponseWriter
	IsIndented() bool
}

type Operations interface {
	Append(...Operation)
	Set(...Operation)
	Get() []Operation
	GetByIndex(int) Operation
	GetByName(string) Operation
	Len() int
}

type ModelStore interface {
	Get()
	Create()
	Edit()
	Remove()
}

type Operation interface {
	BuildPather
	Route() Route
	GetName() string
	GetMethods() []string
	GetPath() string
	GetFullPath() string
	// Param(string) Param
	// Params(...string) Params
	// BodyParam(string) Param
	// BodyParams(...string) Params
	// HeaderParam(string) Param
	// HeaderParams(...string) Params
	// PathParam(string) Param
	// PathParams(...string) Params
	// QueryParam(string) Param
	// QueryParams(...string) Params
}

type Resource interface {
	Create(Context) interface{}
	Retrieve(Context) interface{}
	Update(Context) interface{}
	Delete(Context) interface{}
}

type Values map[interface{}]interface{}

func (vs *Values) Set(k interface{}, v interface{}) {
	*vs[k] = v
}

func (vs *Values) Get(k interface{}) Value {
	return &value{*vs[k]}
}

type Value interface {
	Value() interface{}
	Int() int
	Int64() int64
	Float() float32
	Float64() float64
	String() string
	Bool() bool
}

type Params []Param

type Param interface {
}

type ParamValues map[string]ParamValue

type ParamValue interface {
	Value
	Name() string
	As() ValueType
	RawString() string
	IntE() (int, error)
	Int64E() (int64, error)
	FloatE() (float32, error)
	Float64E() (float64, error)
	BoolE() (bool, error)
}

type BuildPather interface {
	BuildPath(...interface{}) string
}

type Route interface {
	BuildPather
	Path() string
	FullPath() string
	SetPath(string)
}
