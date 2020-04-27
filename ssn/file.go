package ssn

import (
	"os"

	"github.com/gorilla/sessions"
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
