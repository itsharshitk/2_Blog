package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello there!!"})
	})
}
