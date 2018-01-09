package entities

import "gopkg.in/mgo.v2/bson"

//Menu is schema collection
type Menu struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
	Cost int           `json:"cost"`
}

//Places is schema collection
type Place struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name"`
	City      string        `json:"city"`
	Rating    int           `json:"rating"`
	Menu      []Menu        `json:"menu"`
	UpdatedAt int64         `json:"updatedAt" bson:"updatedAt"`
	CreatedAt int64         `json:"createdAt" bson:"createdAt"`
}

//PlaceOut
type PlaceOut struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name"`
	City   string        `json:"city"`
	Rating int           `json:"rating"`
}
