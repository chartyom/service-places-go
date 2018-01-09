package places

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/HenkCord/GOServicePlaces/entities"
	"github.com/HenkCord/GOServicePlaces/repository"
)

type MongoPlacesRepository struct {
	MongoAddr  string
	DB         string
	Collection string
}

func (r *MongoPlacesRepository) connect() (*mgo.Session, *mgo.Collection) {
	s, err := mgo.Dial(r.MongoAddr)
	if err != nil {
		panic(err)
	}
	return s, s.DB(r.DB).C(r.Collection)
}

// Fetch Fetch
func (r *MongoPlacesRepository) Fetch(offset int, limit int, ratingSort string, cityFilter string) []*entities.PlaceOut {
	db, collection := r.connect()
	defer db.Close()
	res := []*entities.PlaceOut{}

	search := bson.M{}

	if cityFilter != "" {
		search = bson.M{
			"city": bson.M{
				"$regex": bson.RegEx{
					Pattern: cityFilter + ".*",
					Options: "i",
				},
			},
		}
	}

	rating := "-createdAt"
	switch ratingSort {
	case "desc":
		rating = "-rating"
	case "asc":
		rating = "rating"
	}

	err := collection.Find(search).Sort(rating).Skip(offset).Limit(limit).All(&res)
	if err != nil {
		//Ошибка может быть как из за отсутствия в бд записи, так и в случае отсутствия подключения к бд
		if err != mgo.ErrNotFound {
			panic(err)
		}
		return res
	}
	return res
}

// GetOne by id
func (r *MongoPlacesRepository) GetOne(id bson.ObjectId) (*entities.PlaceOut, error) {
	db, collection := r.connect()
	defer db.Close()

	res := &entities.PlaceOut{}

	err := collection.Find(bson.M{"_id": id}).One(&res)
	if err != nil {
		//Ошибка может быть как из за отсутствия в бд записи, так и в случае отсутствия подключения к бд
		if err != mgo.ErrNotFound {
			panic(err)
		}
		return nil, errors.New(entities.ErrNotFound)
	}
	return res, nil
}

// Create - create entity
func (r *MongoPlacesRepository) Create(item *entities.Place) (*entities.Place, error) {
	db, collection := r.connect()
	defer db.Close()

	placeExist := &entities.Place{}

	//Проверка существования записи с указанным полем "name"
	err := collection.Find(bson.M{"_id": item.Id}).One(&placeExist)
	if err != nil {
		//Ошибка может быть как из за отсутствия в бд записи, так и в случае подключения к бд
		if err != mgo.ErrNotFound {
			panic(err)
		}
	} else {
		return nil, errors.New(entities.ErrEntryExist)
	}

	item.Id = bson.NewObjectId()

	if len(item.Menu) > 0 {
		for i, _ := range item.Menu {
			item.Menu[i].Id = bson.NewObjectId()
		}
	}

	item.CreatedAt = time.Now().Unix()

	if err = collection.Insert(item); err != nil {
		panic(err)
	}

	return item, nil
}

// Delete Delete
func (r *MongoPlacesRepository) Delete(id bson.ObjectId) (bool, error) {
	db, collection := r.connect()
	defer db.Close()

	if err := collection.Remove(bson.M{"_id": id}); err != nil {
		//Ошибка может быть как из за отсутствия в бд записи, так и в случае подключения к бд
		if err != mgo.ErrNotFound {
			panic(err)
		}
		return false, errors.New(entities.ErrNotFound)
	}
	return true, nil
}

// Update - update entity by id
func (r *MongoPlacesRepository) Update(item *entities.Place) (*entities.Place, error) {
	db, collection := r.connect()
	defer db.Close()

	res := &entities.Place{}

	if err := collection.Find(bson.M{"_id": item.Id}).One(&res); err != nil {
		//Ошибка может быть как из за отсутствия в бд записи, так и в случае подключения к бд
		if err != mgo.ErrNotFound {
			panic(err)
		}
		return nil, errors.New(entities.ErrNotFound)
	}

	if item.Name != "" {
		res.Name = item.Name
	}

	if item.City != "" {
		res.City = item.City
	}

	if item.Rating > 0 {
		res.Rating = item.Rating
	}

	if len(item.Menu) > 0 {
		for i, _ := range item.Menu {
			item.Menu[i].Id = bson.NewObjectId()
		}
		res.Menu = item.Menu
	}

	res.UpdatedAt = time.Now().Unix()
	if err := collection.Update(bson.M{"_id": item.Id}, bson.M{"$set": res}); err != nil {
		panic(err)
	}
	return res, nil
}

func InitMongoPlacesRepository(addr string, db string, collection string) repository.PlacesRepository {
	return &MongoPlacesRepository{
		MongoAddr:  addr,
		DB:         db,
		Collection: collection,
	}
}
