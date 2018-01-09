package places

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/HenkCord/GOServicePlaces/entities"
	"github.com/gin-gonic/gin"
)

//GetMenuByPlaceId
//Params:
//	id objectId
func (s *Server) GetMenuByPlaceId(c *gin.Context) {
	objectId := bson.ObjectIdHex(c.Param("id"))
	list, err := s.PlacesUsecase.GetMenuByPlaceId(objectId)
	if err != nil {
		if err != errors.New(entities.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menu": list})
}
