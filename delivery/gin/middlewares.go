package gin

import (
	"fmt"
	"net/http"

	"github.com/HenkCord/GOServicePlaces/entities"
	"github.com/gin-gonic/gin"
)

//HandlerError middleware
func HandlerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, &gin.H{
					"error":       entities.ErrInternal,
					"description": fmt.Sprint(r),
				})
				return
			}
		}()

		// before request

		c.Next()

		// after request

	}
}
