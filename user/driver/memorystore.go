package driver

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
)

type memoryStore struct {
	data map[string]user.User
}

// NewMemoryStore for users.
func NewMemoryStore() user.Store {
	store := memoryStore{
		data: make(map[string]user.User),
	}
	return &store
}

func (m *memoryStore) FindByMail(mail string) (user.User, error) {
	u, ok := m.data[mail]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	return u, nil
}

func (m *memoryStore) FindAll() (users []user.User, err error) {
	for _, u := range m.data {
		users = append(users, u)
	}
	return users, nil
}

func (m *memoryStore) FindByRegistrationID(registrationID string) (user.User, error) {
	for _, u := range m.data {
		if u.RegistrationID == registrationID {
			return u, nil
		}
	}
	return user.User{}, user.ErrNotFound
}

func (m *memoryStore) Save(u user.User) error {
	m.data[u.Login.Mail] = u
	return nil
}

func (m *memoryStore) RemoveByMail(mail string) error {
	delete(m.data, mail)
	return nil
}
