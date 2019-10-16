package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetHealth returns the healthcheck
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "Health Status: OK!")
}
