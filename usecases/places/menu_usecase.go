package places

import (
	"github.com/HenkCord/GOServicePlaces/entities"
	"gopkg.in/mgo.v2/bson"
)

func (u *PlacesUsecase) GetMenuByPlaceId(id bson.ObjectId) ([]entities.Menu, error) {
	return u.PlacesRepository.GetMenuByPlaceId(id)
}
