package places

import (
	"errors"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/HenkCord/GOServicePlaces/entities"
	"github.com/HenkCord/GOServicePlaces/usecases"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	PlacesUsecase usecases.PlacesUsecase
}

//Fetch
//Query:
//	offset	int	(Default: 0)
//	limit	int	(Default: 30)
//	ratingSort	string	(Default: nil -> asc, desc)
// 	cityFilter	string	(Default: nil -> "search city")
func (s *Server) Fetch(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))
	ratingSort := c.DefaultQuery("rating", "") // asc, desc
	cityFilter := c.DefaultQuery("city", "")

	list, err := s.PlacesUsecase.Fetch(offset, limit, ratingSort, cityFilter)
	if err != nil {
		if err != errors.New(entities.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"places": list})
}

//GetOne
//Params:
//	id objectId
func (s *Server) GetOne(c *gin.Context) {
	objectId := bson.ObjectIdHex(c.Param("id"))
	item, err := s.PlacesUsecase.GetOne(objectId)
	if err != nil {
		if err != errors.New(entities.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

//Update
//Params:
//	id objectId
//Body:
//	name 	string
//	rating	string
//	city	string
//	menu	[]struct
func (s *Server) Update(c *gin.Context) {
	var item entities.Place

	if err := c.ShouldBindWith(&item, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": entities.ErrInvalidArguments})
		return
	}

	item.Id = bson.ObjectIdHex(c.Param("id"))

	res, err := s.PlacesUsecase.Update(&item)
	if err != nil {
		if err != errors.New(entities.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

//Delete
//Params:
//	id objectId
func (s *Server) Delete(c *gin.Context) {
	objectId := bson.ObjectIdHex(c.Param("id"))
	_, err := s.PlacesUsecase.Delete(objectId)
	if err != nil {
		if err != errors.New(entities.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"placeId": objectId})
}

//Create
//Body:
//	name 	string
//	rating	string
//	city	string
//	menu	[]struct
func (s *Server) Create(c *gin.Context) {
	var item entities.Place

	if err := c.ShouldBindWith(&item, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.Id = bson.NewObjectId()

	res, err := s.PlacesUsecase.Create(&item)
	if err != nil {
		if err != errors.New(entities.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func InitPlacesServer(u usecases.PlacesUsecase) *Server {
	return &Server{PlacesUsecase: u}
}
