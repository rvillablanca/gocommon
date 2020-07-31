package ssn

import (
	"github.com/gorilla/sessions"
)

const (
	authKey = "authenticated"
)

var _ Sessioner = &DefaultSessioner{}

type DefaultSessioner struct {
	store       sessions.Store
	sessionName string
}

func (d *DefaultSessioner) Session(r *Req) (*sessions.Session, error) {
	s, err := d.store.Get(r, d.sessionName)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (d *DefaultSessioner) IsAuthenticated(r *Req) (bool, error) {
	s, err := d.Session(r)
	if err != nil {
		return false, err
	}

	return s.Values[authKey] != nil, nil
}

func (d *DefaultSessioner) ClearSession(w Resp, r *Req) error {
	s, err := d.Session(r)
	if err != nil {
		return err
	}

	s.Values = nil
	return d.SaveSession(w, r)
}

func (d *DefaultSessioner) Authenticated(r *Req) {
	_, _ = d.Put(authKey, true, r)
}

func (d *DefaultSessioner) Put(key string, v interface{}, r *Req) (*sessions.Session, error) {
	s, err := d.Session(r)
	if err != nil {
		return nil, err
	}

	s.Values[key] = v
	return s, err
}

func (d *DefaultSessioner) SaveThisSession(s *sessions.Session, w Resp, r *Req) error {
	return s.Save(r, w)
}

func (d *DefaultSessioner) SaveSession(w Resp, r *Req) error {
	s, err := d.Session(r)
	if err != nil {
		return err
	}

	err = s.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (d *DefaultSessioner) DeleteSession(w Resp, r *Req) error {
	s, err := d.Session(r)
	if err != nil {
		return err
	}

	s.Options.MaxAge = -1
	return d.SaveThisSession(s, w, r)
}
