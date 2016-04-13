package framework

import (
	"net/http"
)

type Sessioner interface {
	Sessions() Sessions
	Session(string) Session
	SetSessionValue(string, interface{}, interface{}) error
	GetSessionValue(string, interface{}) (Value, error)
}

type SessionStore interface {
	Names() []string
	Get(*http.Request, string) (Session, error)
	GetAll(*http.Request) (Sessions, error)
	GetMany(*http.Request, ...string) (Sessions, error)
	New(*http.Request, string) (Session, error)
	Save(*http.Request, http.ResponseWriter, Session) error
}

type Sessions []Session

func (s Sessions) Get(name string) Session {
	for _, se := range s {
		if se.Name() == name {
			return se
		}
	}
	return nil
}

func (s *Sessions) Append(nses ...Session) {
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

func (s Sessions) Save(r *http.Request, w http.ResponseWriter) error {
	var err error
	for _, se := range s {
		if !se.Changed() {
			continue
		}
		err = se.Save(r, w)
		if err != nil {
			return err
		}
	}
	return nil
}

type Session interface {
	Set(interface{}, interface{})
	Get(interface{}) Value
	GetOr(interface{}, interface{}) Value
	Unset(interface{}) bool
	ID() string
	Flashes(vars ...string) []Value
	AddFlash(value interface{}, vars ...string)
	Save(*http.Request, http.ResponseWriter) error
	Name() string
	Store() SessionStore
	Values() Values
	Changed() bool
}
