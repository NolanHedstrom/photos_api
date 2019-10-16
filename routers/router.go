package routers

import (
	"photos_api/app/controllers"

	"github.com/gin-gonic/gin"
)

//InitRouters creates the router
func InitRouters(router *gin.Engine) *gin.Engine {

	router.GET("/health", controllers.GetHealth)
	router.GET("/photo", controllers.LoadPhoto)
	router.GET("/photos", controllers.LoadPhotos)

	// v1 := router.Group("/api/v1")
	{
		// v1.GET()
	}

	return router
}
