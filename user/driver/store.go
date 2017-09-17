package driver

import (
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoStore struct {
	session *mgo.Session
	users   *mgo.Collection
}

func NewMongoStore() user.Store {
	databaseName := database.GetEnv("DB_NAME", "wwmbb-dev")
	var store mongoStore

	store.session = database.NewMongoSession()
	store.session.SetMode(mgo.Monotonic, true)
	store.users = store.session.DB(databaseName).C("users")
	return &store
}

func (s *mongoStore) FindByMail(email string) (user.User, error) {
	var findByUserMail user.User
	err := s.users.Find(bson.M{"login.mail": email}).One(&findByUserMail)
	if err == mgo.ErrNotFound {
		return user.User{}, user.ErrNotFound
	}
	return findByUserMail, err
}

func (s *mongoStore) FindByRegistrationID(id string) (user.User, error) {
	var findByRegId user.User
	err := s.users.Find(bson.M{"registrationid": id}).One(&findByRegId)
	if err == mgo.ErrNotFound {
		return user.User{}, user.ErrNotFound
	}
	return findByRegId, err
}

// Returns all _active_ users. This func does not return users which have not finished registering
func (s *mongoStore) FindAll() ([]user.User, error) {
	var results []user.User
	err := s.users.Find(bson.M{"registrationid": ""}).All(&results)
	return results, err
}

// Returns all inactive users --> Users which have not completed the registration process.
func (s *mongoStore) FindAllUnregistered() ([]user.User, error) {
	var results []user.User
	err := s.users.Find(bson.M{"registrationid": bson.M{"$ne": ""}}).All(&results)
	return results, err
}

func (s *mongoStore) Save(user user.User) error {
	_, err := s.users.Upsert(bson.M{"login.mail": user.Mail}, &user)
	return err
}