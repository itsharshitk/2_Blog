package util

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itsharshitk/2_Blog/model"
	"golang.org/x/crypto/bcrypt"
)

var Validate *validator.Validate

func InitValidations() {
	Validate = validator.New()

	Validate.RegisterValidation("validPass", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 6 || len(password) > 20 {
			return false
		}

		upper := regexp.MustCompile(`[A-Z]`)
		lower := regexp.MustCompile(`[a-z]`)
		number := regexp.MustCompile(`[0-9]`)
		symbol := regexp.MustCompile(`[\W_]`)

		return upper.MatchString(password) && lower.MatchString(password) && number.MatchString(password) && symbol.MatchString(password)
	})
}

func ValidateMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "max":
		return err.Field() + " is too long"
	case "min":
		return err.Field() + " is too short"
	case "email":
		return "Incorrect email format"
	case "validPass":
		return "Password must be 6 to 20 characters long, include at least one uppercase letter, one lowercase letter, one number, and one special character"
	default:
		return err.Field() + " is invalid"
	}
}

func HashPass(rawPass string) string {

	if len(rawPass) < 1 {
		log.Fatal("Please enter valid password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(rawPass), 14)

	if err != nil {
		log.Fatal("Password encrytion failed: ", err.Error())
	}
	return string(hash)
}

func GenerateJWTToken(user model.User) (string, error) {
	claims := model.JWTClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "2_Blog",
			Subject:   fmt.Sprintf("%d", user.ID), // converting uint to string
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRETKEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT secret key environment variable (SECRETKEY) not set")
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, err
}
