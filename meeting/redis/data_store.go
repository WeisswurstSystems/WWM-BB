package redis

import (
	"encoding/json"

	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/garyburd/redigo/redis"
)

type Store interface {
	Close()
	Find(id string) (meeting.Meeting, error)
	Save(meeting.Meeting) error
	Delete(id string) error
	FindAll() ([]meeting.Meeting, error)
}

type store struct {
	rc redis.Conn
}

func (store store) Find(id string) (m meeting.Meeting, e error) {
	reply, err := redis.Bytes(store.rc.Do("HGET", "Meetings", id))
	if err != nil {
		return m, err
	}
	return m, json.Unmarshal(reply, &m)
}
func (store store) FindAll() (meetings []meeting.Meeting, e error) {
	reply, err := redis.ByteSlices(store.rc.Do("HVALS", "Meetings"))
	if err != nil {
		return meetings, err
	}
	var meeting meeting.Meeting
	for _, text := range reply {
		err := json.Unmarshal(text, &meeting)
		if err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}
	return meetings, nil
}

func (store store) Save(m meeting.Meeting) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	store.rc.Do("HSET", "Meetings", m.ID, data)

	return nil
}

func (store store) Delete(id string) error {
	_, err := redis.Bytes(store.rc.Do("HDEL", "Meetings", id))
	return err
}

func NewStore(redisURL string) Store {
	rc, err := redis.DialURL(redisURL)
	if err != nil {
		panic(err)
	}
	return store{rc}
}

func (store store) Close() {
	store.rc.Close()
}
