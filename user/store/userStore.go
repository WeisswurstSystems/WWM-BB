package store

import (
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"gopkg.in/mgo.v2/bson"
)

type Store interface {
	Has(email string) (bool, error)
	FindByMail(email string) (user.User, error)
	FindByRegID(id string) (user.User, error)
	FindAll() ([]user.User, error)
	Save(user user.User) (user.User, error)
	Update(user user.User) error
}

func Has(email string) (bool, error) {
	count, err := database.Users.Find(bson.M{"mail": email}).Count()
	return count != 0, err
}

func FindByMail(email string) (user.User, error) {
	var findByUserMail user.User
	err := database.Users.Find(bson.M{"mail": email}).One(&findByUserMail)
	return findByUserMail, err
}

func FindByRegID(id string) (user.User, error) {
	var findByRegId user.User
	err := database.Users.Find(bson.M{"registrationid": id}).One(&findByRegId)
	return findByRegId, err
}

// Returns all _active_ users. This func does not return users which have not finished registering
func FindAll() ([]user.User, error) {
	var results []user.User
	err := database.Users.Find(bson.M{"registrationid": ""}).All(&results)
	return results, err
}

// Returns all inactive users --> Users which have not completed the registration process.
func FindAllUnregistered() ([]user.User, error) {
	var results []user.User
	err := database.Users.Find(bson.M{"registrationid": bson.M{"$ne" : ""}}).All(&results)
	return results, err
}

func Save(user user.User) (user.User, error) {
	err := database.Users.Insert(&user)
	return user, err
}

func Update(user user.User) error {
	return database.Users.Update(bson.M{"mail": user.Mail}, user)
}
