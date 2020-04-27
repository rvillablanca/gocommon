package ssn

import "github.com/boj/redistore"

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
