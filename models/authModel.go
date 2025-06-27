package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null" json:"username" validate:"required,max=50,min=2"`
	Email     string    `gorm:"not null" json:"email" validate:"required,email"`
	Password  string    `gorm:"not null" json:"password" validate:"required,validPass"`
	Role      string    `gorm:"default:'reader'"` //reader, admin, editor
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"-"`
	Posts     []Post    `gorm:"foreignKey:UserId"`
	Comments  []Comment
	Likes     []Like
}
