package route

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/controller"
	"github.com/itsharshitk/2_Blog/middleware"
)

func GetRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/signup", controller.SignUp)
	r.POST("/refresh", controller.RefreshHandler)

	p := r.Group("/", middleware.JWTMiddleware())

	p.POST("/logout", controller.Logout)

	p.POST("/post", controller.CreatePost)

}
