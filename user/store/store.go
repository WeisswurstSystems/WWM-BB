package store

import (
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"gopkg.in/mgo.v2/bson"
)

type Store interface {
	Has(email string) (bool, error)
	FindByMail(email string) (user.User, error)
	FindAll() ([]user.User, error)
	Save(user user.User) error
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

func FindAll() ([]user.User, error) {
	var results []user.User
	err := database.Users.Find(nil).All(&results)
	return results, err
}

func Save(user user.User) error {
	return database.Users.Insert(&user)
}
