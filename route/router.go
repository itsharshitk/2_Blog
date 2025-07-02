package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/controller"
	"github.com/itsharshitk/2_Blog/middleware"
)

func GetRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/signup", controller.SignUp)
	r.POST("/refresh", controller.RefreshHandler)

	p := r.Group("/api/v1", middleware.JWTMiddleware())

	p.GET("/co", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello there"})
	})
}
