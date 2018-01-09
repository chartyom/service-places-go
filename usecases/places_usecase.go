package usecases

import (
	"github.com/HenkCord/GOServicePlaces/entities"
	"gopkg.in/mgo.v2/bson"
)

type PlacesUsecase interface {
	Fetch(offset int, limit int, ratingSort string, cityFilter string) ([]*entities.PlaceOut, error)
	GetOne(id bson.ObjectId) (*entities.PlaceOut, error)
	Update(item *entities.Place) (*entities.Place, error)
	Create(item *entities.Place) (*entities.Place, error)
	Delete(id bson.ObjectId) (bool, error)
	GetMenuByPlaceId(id bson.ObjectId) ([]entities.Menu, error)
}
