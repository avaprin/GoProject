package product

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repo interface {
	Create(cloth *Cloth) error
	Get(id bson.ObjectId) (Cloth, error)
	GetCategory(category string) ([]Cloth, error)
	Update(cloth *Cloth) error
	Delete(id bson.ObjectId) error
}

type repo struct {
}

func (r repo) Create(product *Cloth) error {
	var s, collection, df = Dial()
	if df != nil {
		return df
	}
	product.Id = bson.NewObjectId()
	var err = collection.Insert(&product)
	s.Close()
	return err
}

func (r repo) Get(id bson.ObjectId) (Cloth, error) {
	var s, collection, df = Dial()
	if df != nil {
		return Cloth{}, df
	}
	var product = Cloth{}
	var err = collection.FindId(id).One(&product)
	s.Close()
	return product, err
}

func (r repo) GetCategory(category string) ([]Cloth, error) {
	var s, collection, df = Dial()
	if df != nil {
		return []Cloth{}, df
	}
	var products = []Cloth{}
	var err = collection.Find(bson.M{"category": category}).All(&products)
	s.Close()
	return products, err
}

func (r repo) Update(product *Cloth) error {
	var s, collection, df = Dial()
	if df != nil {
		return df
	}
	var dbProduct, errNotFound = r.Get(product.Id)
	if errNotFound != nil {
		return errNotFound
	}

	var err = collection.Update(dbProduct, product)
	s.Close()
	return err
}

func (r repo) Delete(id bson.ObjectId) error {
	var s, collection, df = Dial()
	if df != nil {
		return df
	}
	var err = collection.RemoveId(id)
	s.Close()
	return err
}

func NewRepo() Repo {
	return repo{}
}

func Dial() (*mgo.Session, *mgo.Collection, error) {
	var s, err = mgo.Dial("localhost:27017")
	if err == nil {
		var collection = s.DB("ClothingShop").C("Products")
		return s, collection, nil
	}
	return nil, nil, errors.New("dial failed")
}