package redis

import (
	"github.com/garyburd/redigo/redis"
)

type EventStore interface {
}

type eventStore struct {
	rc redis.Conn
}

func NewEventStore(redisURL string) EventStore {
	rc, err := redis.DialURL(redisURL)
	if err != nil {
		panic(err)
	}
	return eventStore{rc}
}

func (store eventStore) Close() {
	store.rc.Close()
}
