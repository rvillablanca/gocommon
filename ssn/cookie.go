package ssn

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func InitCookieSession(maxAge int, name, path string) {
	sessionName = name
	cookieSession := sessions.NewCookieStore(securecookie.GenerateRandomKey(10))
	cookieSession.MaxAge(maxAge)
	cookieSession.Options.Path = path

	store = cookieSession
}
