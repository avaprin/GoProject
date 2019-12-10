package user

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	Create(user *User) error
	Get(user *User) (*User, error)
	Delete(login string) error
	SaveUserChange(login string, changedTo *User) error
}

type repo struct {
}

func (r repo) Create(user *User) error {
	var s, collection, df = Dial()
	if df != nil {
		return df
	}
	var c, _ = collection.Find(bson.M{"email": user.Email}).Count()
	if c != 0 {
		return errors.New("user not unique")
	}
	user.Id = bson.NewObjectId()
	var err = collection.Insert(&user)
	s.Close()
	return err
}

func (r repo) Get(user *User) (*User, error) {
	var s, collection, df = Dial()
	if df != nil {
		return nil, df
	}
	var userFromDB = User{}
	var err = collection.Find(bson.M{"email": user.Email}).One(&userFromDB)
	s.Close()
	return &userFromDB, err
}

func (r repo) Delete(login string) error {
	var s, collection, df = Dial()
	if df != nil {
		return df
	}
	var err = collection.Remove(bson.M{"email": login})
	s.Close()
	return err
}

func (r repo) SaveUserChange(login string, update *User) error {
	var s, collection, df = Dial()
	if df != nil {
		return df
	}
	var userFromDB = User{}
	var err = collection.Find(bson.M{"email": login}).One(&userFromDB)
	if err != nil {
		return err
	}

	update.Id = userFromDB.Id
	var response = collection.Update(userFromDB, update)
	s.Close()
	return response
}

func NewUserRepo() Repository {
	return repo{}
}

func Dial() (*mgo.Session, *mgo.Collection, error) {
	var s, err = mgo.Dial("localhost:27017")
	if err == nil {
		var collection = s.DB("ClothingShop").C("Users")
		return s, collection, nil
	}
	return nil, nil, errors.New("dial failed")
}
