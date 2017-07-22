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

func (s *mongoStore) HasByMail(email string) (bool, error) {
	count, err := s.users.Find(bson.M{"mail": email}).Count()
	return count != 0, err
}

func (s *mongoStore) FindByMail(email string) (user.User, error) {
	var findByUserMail user.User
	err := s.users.Find(bson.M{"mail": email}).One(&findByUserMail)
	return findByUserMail, err
}

func (s *mongoStore) FindByRegistrationID(id string) (user.User, error) {
	var findByRegId user.User
	err := s.users.Find(bson.M{"registrationid": id}).One(&findByRegId)
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
	err := s.users.Insert(&user)
	return err
}

func (s *mongoStore) Update(user user.User) error {
	return s.users.Update(bson.M{"mail": user.Mail}, user)
}
