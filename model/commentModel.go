package model

import "time"

type Comment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserId      uint      `json:"user_id" validate:"required"`
	PostId      uint      `json:"post_id" validate:"required"`
	CommentText string    `gorm:"type:text;not null" json:"comment_text" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `gorm:"index" json:"-"`
}
