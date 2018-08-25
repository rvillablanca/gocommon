// Package cookiesession is a cookie based session store, so it is not intended to save sensible data.
// This package has two main purposes:
//   1. Serve as session store
//   2. Identified if current user is authenticated.
// Some methods panic because getting session or saving values are critical and failing is considered a non recoverable error.
package cookiesession

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const (
	authKey = "authenticated"
)

var (
	defaultCookieName = "cookiesession"
	store             *sessions.CookieStore
)

// InitSession initialize the package with the name of the cookie
func InitSession(cookieName string, maxAge int, path string) {
	store = sessions.NewCookieStore(securecookie.GenerateRandomKey(10))
	store.MaxAge(maxAge)
	store.Options.Path = path
	defaultCookieName = cookieName
}

// Session allows get session reference
func Session(r *http.Request) *sessions.Session {
	s, err := store.Get(r, defaultCookieName)
	if err != nil {
		//If this ocurrs, then panic will be recovered ant 500 error page will be rendered
		panic("could not get session object")
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
		panic("could not save session object")
	}
}
