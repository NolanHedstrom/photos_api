package main

import (
	"photos_api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router = routers.InitRouters(router)

	router.Run(":8080")
}
