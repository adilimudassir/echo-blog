package models

import "gorm.io/gorm"

// Post model represents a blog post
type Post struct {
    gorm.Model
    Title   string   `json:"title" validate:"required"`
    Content string   `json:"content" validate:"required"`
    Comments []Comment `gorm:"foreignKey:PostID"`
}
