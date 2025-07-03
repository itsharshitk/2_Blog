package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/model"
)

func CreatePost(c *gin.Context) {

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Resp{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
		return
	}

	post := model.Post{
		UserId:  c.GetUint("UserId"),
		Title:   req.Title,
		Content: req.Content,
		Slug: ,
	}

	result := db.DB.Create(&post)
}
