package ssn

import (
	"github.com/boj/redistore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"os"
)

func NewCookieSessioner(maxAge int, sessionName, path string) *DefaultSessioner {
	cookieSession := sessions.NewCookieStore(securecookie.GenerateRandomKey(10))
	cookieSession.MaxAge(maxAge)
	cookieSession.Options.Path = path
	return &DefaultSessioner{store: cookieSession, sessionName: sessionName}
}

func NewFileSessioner(maxAge int, path, sessionName, key, contextPath string) (*DefaultSessioner, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0755); err != nil {
			return nil, err
		}

	}
	fileSystemStore := sessions.NewFilesystemStore(path, []byte(key))
	fileSystemStore.MaxAge(maxAge)
	fileSystemStore.Options.Path = contextPath

	return &DefaultSessioner{store: fileSystemStore, sessionName: sessionName}, nil
}

func NewRedisSessioner(maxAgeStore int, sessionName, address, password, key string) (*DefaultSessioner, error) {
	redisStore, err := redistore.NewRediStore(10, "tcp", address, password, []byte(key))
	if err != nil {
		return nil, err
	}

	redisStore.SetMaxAge(maxAgeStore)
	return &DefaultSessioner{store: redisStore, sessionName: sessionName}, nil
}
