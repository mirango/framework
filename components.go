package framework

type Locale interface {
	Code() string
}
type User interface {
	Locale() string
}
