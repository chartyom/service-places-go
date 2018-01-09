package repository

import (
	"github.com/HenkCord/GOServicePlaces/entities"
	"gopkg.in/mgo.v2/bson"
)

// PlacesRepository - interface repository
type PlacesRepository interface {
	Fetch(offset int, limit int, ratingSort string, cityFilter string) []*entities.PlaceOut
	GetOne(id bson.ObjectId) (*entities.PlaceOut, error)
	Update(item *entities.Place) (*entities.Place, error)
	Create(item *entities.Place) (*entities.Place, error)
	Delete(id bson.ObjectId) (bool, error)
	GetMenuByPlaceId(id bson.ObjectId) ([]entities.Menu, error)
}
