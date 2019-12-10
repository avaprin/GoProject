package product

import "gopkg.in/mgo.v2/bson"

type Cloth struct {
	Id           bson.ObjectId `bson:"_id"`
	Category     string
	Manufacturer string
	Name         string
	Season       string
	Size         string
	Collection   string
	Price        float64
	Text         string
}
