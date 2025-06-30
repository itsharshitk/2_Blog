package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/itsharshitk/2_Blog/db"
	"github.com/itsharshitk/2_Blog/model"
	"github.com/itsharshitk/2_Blog/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var req model.LoginRequest
	var foundUser model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Resp{
			Status:       http.StatusBadRequest,
			Message:      "Invalid input",
			ErrorDetails: err.Error(),
		})
		return
	}

	if err := util.Validate.Struct(req); err != nil {
		errs := make(map[string]string)
		for _, val := range err.(validator.ValidationErrors) {
			errs[val.Field()] = util.ValidateMessage(val)
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": errs})
		return
	}

	if err := db.DB.Where("email = ?", req.Email).First(&foundUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, model.Resp{
				Status:  http.StatusUnauthorized,
				Message: "User Not Registered",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Database error occurred",
			ErrorDetails: err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, model.Resp{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Password",
		})
		return
	}

	token, err := util.GenerateJWTToken(foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Token Generation Failed",
			ErrorDetails: err,
		})
	}

	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "Login successful",
		Data: model.UserResponse{
			ID:       foundUser.ID,
			Username: foundUser.Username,
			Email:    foundUser.Email,
			Role:     foundUser.Role,
			Token:    token,
		},
	})

}

func SignUp(c *gin.Context) {
	var req model.SignUpRequest
	var foundUser model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Resp{
			Status:       http.StatusBadRequest,
			Message:      "Invalid request",
			ErrorDetails: err.Error(),
		})
		return
	}

	if err := util.Validate.Struct(req); err != nil {
		errs := make(map[string]string)
		for _, val := range err.(validator.ValidationErrors) {
			errs[val.Field()] = util.ValidateMessage(val)
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": errs})
		return
	}

	err := db.DB.Where("email = ?", req.Email).First(&foundUser).Error
	if err == nil {
		c.JSON(http.StatusConflict, model.Resp{
			Status:  http.StatusConflict,
			Message: "User already exists",
		})
		return
	}

	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "Something went wrong!",
			ErrorDetails: err.Error(),
		})
		return
	}

	req.Password = util.HashPass(req.Password)

	newUser := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.Resp{
			Status:       http.StatusInternalServerError,
			Message:      "User Not Created",
			ErrorDetails: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Resp{
		Status:  http.StatusOK,
		Message: "User created successfully",
		Data:    newUser,
	})
}
