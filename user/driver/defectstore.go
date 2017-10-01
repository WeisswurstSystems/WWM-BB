package driver

import "github.com/WeisswurstSystems/WWM-BB/user"

type defectStore struct {
	defect error
}

// NewDefectStore for a store that always error with the given error.
func NewDefectStore(defect error) user.Store {
	store := defectStore{defect: defect}
	return &store
}

func (ds *defectStore) FindByMail(string) (user.User, error) {
	return user.User{}, ds.defect
}

func (ds *defectStore) FindAll() ([]user.User, error) {
	return []user.User{}, ds.defect
}

func (ds *defectStore) FindByRegistrationID(registrationID string) (user.User, error) {
	return user.User{}, ds.defect
}

func (ds *defectStore) Save(user.User) error {
	return ds.defect
}

func (ds *defectStore) RemoveByMail(string) error {
	return ds.defect
}
