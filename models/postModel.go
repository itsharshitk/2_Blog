package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id" validate:"required"`
	Title     string    `gorm:"not null" json:"title" validate:"required,min=5"`
	Content   string    `gorm:"type:text;not null" json:"content" validate:"required,min=10"`
	Slug      string    `json:"slug" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"-"`
	Likes     []Like    `gorm:"foreignKey:PostId"`
	Comments  []Comment `gorm:"foreignKey:PostId"`
}
