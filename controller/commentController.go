package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/model"
)

func AddComment(c *gin.Context) {
	var comment struct {
		UserId      uint   `json:"user_id" binding:"required"`
		PostId      uint   `json:"post_id" binding:"required"`
		CommentText string `json:"comment_text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, model.Resp{
			Status:       http.StatusBadRequest,
			Message:      "Invalid Request",
			ErrorDetails: err.Error(),
		})
		return
	}

	var cmt model.Comment
	cmt.UserId = comment.UserId
	cmt.PostId = comment.PostId
	cmt.CommentText = comment.CommentText

	result := db.DB.Create(&cmt)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Someting went wrong while saving comment",
			ErrorDetails: result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "Comment saved successfully",
	})
}
