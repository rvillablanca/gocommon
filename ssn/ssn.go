package ssn

import (
	"fmt"
	"net/http"
	"os"

	"github.com/boj/redistore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const (
	authKey = "authenticated"
)

var (
	sessionName = "default-ssn"
	store       sessions.Store
)

func InitFileSystemSession(maxAge int, path, name, key, contextPath string) error {
	sessionName = name
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0755); err != nil {
			return err
		}

	}
	fileSystemStore := sessions.NewFilesystemStore(path, []byte(key))
	fileSystemStore.MaxAge(maxAge)
	fileSystemStore.Options.Path = contextPath

	store = fileSystemStore
	return nil
}

func InitCookieSession(maxAge int, name, path string) {
	sessionName = name
	cookieSession := sessions.NewCookieStore(securecookie.GenerateRandomKey(10))
	cookieSession.MaxAge(maxAge)
	cookieSession.Options.Path = path

	store = cookieSession
}

func InitRedisSession(maxAgeStore int, name, address, password, key string) error {
	sessionName = name
	redisStore, err := redistore.NewRediStore(10, "tcp", address, password, []byte(key))
	if err != nil {
		return err
	}

	redisStore.SetMaxAge(maxAgeStore)
	store = redisStore
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
