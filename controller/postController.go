package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/model"
	"github.com/itsharshitk/2_Blog/util"
)

func CreatePost(c *gin.Context) {

	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Resp{
			Status:       http.StatusBadRequest,
			Message:      "Invalid request",
			ErrorDetails: err.Error(),
		})
		return
	}

	slug := util.Slugify(req.Title)

	post := model.Post{
		UserId:  c.GetUint("userId"),
		Title:   req.Title,
		Content: req.Content,
		Slug:    slug,
	}

	result := db.DB.Create(&post)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Failed to save Post",
			ErrorDetails: result.Error.Error(),
		})
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Someting went wrong while saving post",
			ErrorDetails: result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "Post created successfully",
	})
}
