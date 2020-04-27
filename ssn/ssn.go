package ssn

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	authKey = "authenticated"
)

var (
	sessionName = "default-ssn"
	store       sessions.Store
)

// Session allows get session reference
func Session(r *http.Request) *sessions.Session {
	s, err := store.Get(r, sessionName)
	if err != nil {
		fmt.Printf("warning: an error ocurred while getting Session from request: %v", err)
	}
	return s
}

// IsAuthenticated allows to know if the current session is authenticated
func IsAuthenticated(r *http.Request) bool {
	s := Session(r)
	return s.Values[authKey] != nil
}

// ClearSession allows delete all content of the session
func ClearSession(w http.ResponseWriter, r *http.Request) error {
	s := Session(r)
	s.Values = nil
	return SaveSession(w, r)
}

// Authenticated mark current session as authenticated
func Authenticated(w http.ResponseWriter, r *http.Request) {
	Put(authKey, "authenticated", r)
}

// Save allows put a value with the specified key and value
func Put(key string, v interface{}, r *http.Request) *sessions.Session {
	s := Session(r)
	s.Values[key] = v
	return s
}

// SaveThisSession save the provided session
func SaveThisSession(s *sessions.Session, w http.ResponseWriter, r *http.Request) error {
	return s.Save(r, w)
}

// SaveSession allows to save all values in the session reference
func SaveSession(w http.ResponseWriter, r *http.Request) error {
	return Session(r).Save(r, w)
}

// DeleteSession delete the current session
func DeleteSession(w http.ResponseWriter, r *http.Request) error {
	s := Session(r)
	s.Options.MaxAge = -1
	return SaveThisSession(s, w, r)
}
