package gin

import (
	e "github.com/HenkCord/GOServicePlaces/delivery/gin/places"
	"github.com/HenkCord/GOServicePlaces/usecases"
	"github.com/gin-gonic/gin"
)

//InitPlacesServer init
func InitPlacesServer(r *gin.Engine, u usecases.PlacesUsecase) {
	s := e.InitPlacesServer(u)
	places := r.Group("/places")
	{
		places.GET("", s.Fetch)
		places.GET("/:id", s.GetOne)
		places.POST("", s.Create)
		places.PUT("/:id", s.Update)
		places.DELETE("/:id", s.Delete)
		places.GET("/:id/menu", s.GetMenuByPlaceId)
	}
}
