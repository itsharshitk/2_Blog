package model

import "time"

type Like struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id" validate:"required"`
	PostId    uint      `json:"post_id" validate:"required"`
	IsLiked   bool      `json:"is_liked" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}
