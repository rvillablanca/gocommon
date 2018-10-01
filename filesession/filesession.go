package filesession

// Package filesession is a set of utilities on top of gorilla file sessions.

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

const (
	authKey = "authenticated"
)

var (
	sessionName = "fsid"
	store       *sessions.FilesystemStore
)

// InitSession initialize the package with the name of the cookie
func InitSession(maxAge int, path, name, key, contextPath string) error {
	sessionName = name
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0755); err != nil {
			return err
		}

	}
	store = sessions.NewFilesystemStore(path, []byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = contextPath
	return nil
}

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
func ClearSession(w http.ResponseWriter, r *http.Request) {
	s := Session(r)
	s.Values = nil
	SaveSession(w, r)
}

// Authenticated mark current session as authenticated
func Authenticated(w http.ResponseWriter, r *http.Request) {
	s := Save(authKey, "authenticated", r)
	SaveThisSession(s, w, r)
}

// Save allows put a value with the specified key and value
func Save(key string, v interface{}, r *http.Request) *sessions.Session {
	s := Session(r)
	s.Values[key] = v
	return s
}

// SaveThisSession save the provided session
func SaveThisSession(s *sessions.Session, w http.ResponseWriter, r *http.Request) {
	err := s.Save(r, w)
	if err != nil {
		panic("could not save session object:" + err.Error())
	}
}

// SaveSession allows to save all values in the session reference
func SaveSession(w http.ResponseWriter, r *http.Request) {
	err := Session(r).Save(r, w)
	if err != nil {
		panic("could not save session object:" + err.Error())
	}
}

// DeleteSession delete the current session
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	s := Session(r)
	s.Options.MaxAge = -1
	SaveThisSession(s, w, r)
}
