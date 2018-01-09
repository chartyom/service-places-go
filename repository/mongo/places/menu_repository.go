package places

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/HenkCord/GOServicePlaces/entities"
)

// GetMenuByPlaceId
func (r *MongoPlacesRepository) GetMenuByPlaceId(id bson.ObjectId) ([]entities.Menu, error) {
	db, collection := r.connect()
	defer db.Close()
	res := &entities.Place{}
	err := collection.Find(bson.M{"_id": id}).Select(bson.M{"menu": 1}).One(&res)
	if err != nil {
		//Ошибка может быть как из за отсутствия в бд записи, так и в случае отсутствия подключения к бд
		if err != mgo.ErrNotFound {
			panic(err)
		}
		return nil, errors.New(entities.ErrNotFound)
	}
	return res.Menu, nil
}
