package places

import (
	"errors"

	"github.com/HenkCord/GOServicePlaces/entities"
	"github.com/asaskevich/govalidator"
)

// check required arguments
func fetchValidation(offset *int, limit *int, ratingSort *string) error {

	if *ratingSort != "" && *ratingSort != "asc" && *ratingSort != "desc" {
		*ratingSort = ""
	}

	if *offset < 0 {
		*offset = 0
	}

	if *limit < 1 || *limit > 1000 {
		*limit = 30
	}

	return nil
}

func createValidation(item *entities.Place) error {
	var common *entities.Common
	arr := make(map[string]string)

	if !govalidator.StringLength(item.Name, "2", "60") {
		arr["name"] = entities.ErrInvalidName
	}
	if !govalidator.StringLength(item.City, "2", "60") {
		arr["city"] = entities.ErrInvalidCity
	}
	if item.Rating < 0 || item.Rating > 5 {
		arr["rating"] = entities.ErrInvalidRating
	}

	if len(item.Menu) > 0 {
		for i, _ := range item.Menu {
			if !govalidator.StringLength(item.Menu[i].Name, "2", "60") {
				arr["menu"] = entities.ErrInvalidMenu
			}
			if item.Menu[i].Cost < 0 {
				arr["menu"] = entities.ErrInvalidMenu
			}
		}
	}

	if len(arr) > 0 {
		return errors.New(common.StringError(arr))
	}
	return nil
}

func updateValidation(item *entities.Place) error {
	var common *entities.Common
	arr := make(map[string]string)

	if !govalidator.StringLength(item.Name, "2", "60") {
		arr["name"] = entities.ErrInvalidName
	}
	if !govalidator.StringLength(item.City, "2", "60") {
		arr["city"] = entities.ErrInvalidCity
	}
	if item.Rating < 0 || item.Rating > 5 {
		arr["rating"] = entities.ErrInvalidRating
	}

	if len(item.Menu) > 0 {
		for i, _ := range item.Menu {
			if !govalidator.StringLength(item.Menu[i].Name, "2", "60") {
				arr["menu"] = entities.ErrInvalidMenu
			}
			if item.Menu[i].Cost < 0 {
				arr["menu"] = entities.ErrInvalidMenu
			}
		}
	}

	if len(arr) > 0 {
		return errors.New(common.StringError(arr))
	}
	return nil
}
