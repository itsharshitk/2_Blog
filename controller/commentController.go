package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/model"
)

func AddComment(c *gin.Context) {

	postId, err := strconv.ParseUint(c.Param("postId"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userId := c.GetUint("userId")

	var comment struct {
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
	cmt.UserId = userId
	cmt.PostId = uint(postId)
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

func CommentsOnPost(c *gin.Context) {
	postId := c.Param("postId")

	comments := model.Comment{}
	result := db.DB.Where("post_id = ?", postId).Find(&comments)
}

func UpdateComment(c *gin.Context) {

	CommentId := c.Param("commentId")

	var comment struct {
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

	result := db.DB.Model(&model.Comment{}).Where("id = ?", CommentId).Update("comment_text", comment.CommentText)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Failed to update comment",
			ErrorDetails: result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, model.Resp{
			Status:  http.StatusNotFound,
			Message: "Comment dosen't exist",
		})
		return
	}

	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "Comment updated successfully",
	})
}
