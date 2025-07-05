package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"not null" json:"username" validate:"required,max=50,min=2"`
	Email        string         `gorm:"not null;unique" json:"email" validate:"required,email"`
	Password     string         `gorm:"not null" json:"password" validate:"required,validPass"`
	Role         string         `gorm:"type:enum('reader','admin','editor');default:'reader'" json:"role"`
	UpdatedAt    time.Time      `json:"updated_at"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	RefreshToken []RefreshToken `gorm:"foreignKey:UserId"`
	Posts        []Post         `gorm:"foreignKey:UserId"`
	Comments     []Comment      `gorm:"foreignKey:UserId"`
	Likes        []Like         `gorm:"foreignKey:UserId"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,validPass"`
}

type SignUpRequest struct {
	Username             string `json:"username" binding:"required,min=3,max=30" validate:"required,min=3,max=30"`
	Email                string `json:"email" binding:"required,email" validate:"required,email"`
	Password             string `json:"password" binding:"required" validate:"required,validPass"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required" validate:"required,eqfield=Password"`
}

type UserResponse struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	JWTToken     string `json:"jwt_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `gorm:"index" json:"user_id"`
	Token     string    `gorm:"uniqueIndex" json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
	CreatedAt time.Time `json:"created_at"`
}
