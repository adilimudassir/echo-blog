package models

import "gorm.io/gorm"

// Comment represents a comment on a blog post
type Comment struct {
	gorm.Model

    Author    string    `json:"author" validate:"required"`
    Content   string    `json:"content" validate:"required"`
	
 	PostID    uint
    Post      Post      `gorm:"foreignKey:PostID"`
}
