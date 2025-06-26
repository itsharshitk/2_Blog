package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/routes"
)

func main() {
	r := gin.Default()
	db.ConnectDB()
	routes.GetRoutes(r)

	r.Run(":8080")
}
