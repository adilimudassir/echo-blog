package models

import "gorm.io/gorm"

// Comment represents a comment on a blog post
type Comment struct {
	gorm.Model

    Author    string    `json:"author" validate:"required"`
    Content   string    `json:"content" validate:"required"`
	
 	PostID    uint   	`gorm:"index" json:"post_id" validate:"required"`
    Post      Post      `gorm:"foreignKey:PostID"`

}
