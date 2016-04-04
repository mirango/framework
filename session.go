package framework

import (
	"net/http"
)

type Sessioner interface {
	Sessions() Sessions
	Session(string) (Session, error)
	SetSessionValue(string, interface{}, interface{}) error
	GetSessionValue(string, interface{}) (Value, error)
}

type SessionStore interface {
	SessionNames() []string
	Get(*http.Request, string) (Session, error)
	GetAll(*http.Request) (Sessions, error)
	GetMany(*http.Request, ...string) (Sessions, error)
	New(*http.Request, string) (Session, error)
	Save(Response, Session) error
}

type Sessions []Session

func (s Sessions) Get(name string) (Session, error) {
	for _, se := range s {
		if se.Name() == name {
			return se, nil
		}
	}
	return nil, nil
}

func (s *Sessions) Add(ns Session) {
	if ns == nil {
		return
	}
	for i, se := range *s {
		if se.Name() == ns.Name() {
			(*s)[i] = ns
			return
		}
	}
	*s = append(*s, ns)
}

func (s *Sessions) AddMany(nses ...Session) {
	for _, ns := range nses {
		if ns == nil {
			continue
		}
		for i, se := range *s {
			if se.Name() == ns.Name() {
				(*s)[i] = ns
				continue
			}
		}
		*s = append(*s, ns)
	}
}

func (s Sessions) Save(c Context) error {
	var err error
	for _, se := range s {
		err = se.Save(c)
		if err != nil {
			return err
		}
	}
	return nil
}

type Session interface {
	Flashes(vars ...string) []Value
	AddFlash(value interface{}, vars ...string)
	Save(Context) error
	Name() string
	Store() SessionStore
}
