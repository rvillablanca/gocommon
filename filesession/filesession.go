package filesession

// Package filesession is a set of utilities on top of gorilla file sessions.

import (
	"net/http"

	"github.com/gorilla/securecookie"
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
func InitSession(path string, maxAge int, name string) {
	sessionName = name
	store = sessions.NewFilesystemStore(path, securecookie.GenerateRandomKey(10))
	store.MaxAge(maxAge)
}

// Session allows get session reference
func Session(r *http.Request) *sessions.Session {
	s, err := store.Get(r, sessionName)
	if err != nil {
		//If this ocurrs, then panic will be recovered ant 500 error page will be rendered
		panic("could not get session object:" + err.Error())
	}
	return s
}

// IsAuthenticated allows to know if the current session is authenticated
func IsAuthenticated(r *http.Request) bool {
	s := Session(r)
	return s.Values[authKey] != nil
}

// InvalidateSession allows delete al content of the session
func InvalidateSession(w http.ResponseWriter, r *http.Request) {
	s := Session(r)
	s.Values = nil
	SaveSession(w, r)
}

// Authenticated mark current session as authenticated
func Authenticated(w http.ResponseWriter, r *http.Request) {
	Save(authKey, "authenticated", r)
	SaveSession(w, r)
}

// Save allows put a value with the specified key and value
func Save(key string, v interface{}, r *http.Request) {
	Session(r).Values[key] = v
}

// SaveSession allows to save all values in the session reference
func SaveSession(w http.ResponseWriter, r *http.Request) {
	err := Session(r).Save(r, w)
	if err != nil {
		panic("could not save session object:" + err.Error())
	}
}
