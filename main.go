package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/route"
	"github.com/itsharshitk/2_Blog/util"
)

func main() {
	r := gin.Default()
	db.ConnectDB()
	util.InitValidations()
	route.GetRoutes(r)

	r.Run(":8080")
}
