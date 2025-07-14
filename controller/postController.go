package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/model"
	"github.com/itsharshitk/2_Blog/util"
	"gorm.io/gorm"
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

func GetAllPosts(c *gin.Context) {
	var posts []model.Post
	if result := db.DB.Find(&posts).Error; result != nil {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Something went wrong",
			ErrorDetails: result.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "Fetched all posts",
		Data:    posts,
	})
}

func PostById(c *gin.Context) {
	postId := c.Param("postId")
	var post model.Post

	if result := db.DB.Where("id = ?", postId).First(&post).Error; result != nil {
		if errors.Is(result, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, model.Resp{
				Status:       http.StatusNotFound,
				Message:      "No Such Post Found",
				ErrorDetails: result.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, model.Resp{
				Status:       http.StatusInternalServerError,
				Message:      "Failed to fetch post",
				ErrorDetails: result.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "Fetched post successfuly",
		Data:    post,
	})
}

func DelPost(c *gin.Context) {
	postId := c.Param("postId")
	var post model.Post

	if err := db.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, model.Resp{
				Status:       http.StatusNotFound,
				Message:      "No such post exist or it is already deleted",
				ErrorDetails: err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, model.Resp{
				Status:       http.StatusInternalServerError,
				Message:      "Failed to find post",
				ErrorDetails: err.Error(),
			})
			return
		}
	}

	if err := db.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Failed to delete post",
			ErrorDetails: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "Post deleted successfuly",
	})
}
