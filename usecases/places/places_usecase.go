package places

import (
	"github.com/HenkCord/GOServicePlaces/entities"
	"github.com/HenkCord/GOServicePlaces/repository"
	"github.com/HenkCord/GOServicePlaces/usecases"
	"gopkg.in/mgo.v2/bson"
)

type PlacesUsecase struct {
	PlacesRepository repository.PlacesRepository
}

func (u *PlacesUsecase) GetOne(id bson.ObjectId) (*entities.PlaceOut, error) {
	return u.PlacesRepository.GetOne(id)
}

func (u *PlacesUsecase) Fetch(offset int, limit int, ratingSort string, cityFilter string) ([]*entities.PlaceOut, error) {
	fetchValidation(&offset, &limit, &ratingSort)
	return u.PlacesRepository.Fetch(offset, limit, ratingSort, cityFilter), nil
}

func (u *PlacesUsecase) Delete(id bson.ObjectId) (bool, error) {
	return u.PlacesRepository.Delete(id)
}

// Create Create
func (u *PlacesUsecase) Create(item *entities.Place) (*entities.Place, error) {
	if err := createValidation(item); err != nil {
		return nil, err
	}
	res, err := u.PlacesRepository.Create(item)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *PlacesUsecase) Update(item *entities.Place) (*entities.Place, error) {

	if err := updateValidation(item); err != nil {
		return nil, err
	}

	return u.PlacesRepository.Update(item)
}

func InitPlacesUsecase(u repository.PlacesRepository) usecases.PlacesUsecase {
	return &PlacesUsecase{PlacesRepository: u}
}
