package models

import "gorm.io/gorm"

// Comment represents a comment on a blog post
type Comment struct {
	gorm.Model
    Author    string    `json:"author" validate:"required"`
    Content   string    `json:"content" validate:"required"`
 	PostID    uint      `json:"post_id" validate:"required"`
    Post      Post      `json:"posts" gorm:"foreignKey:PostID" validate:"omitempty"`
}
