
package db

import (
    "github.com/adilimudassir/echo-blog/internal/models"
    "gorm.io/gorm"
)

// PostRepository represents the repository for interacting with posts
type PostRepository struct {
    db *gorm.DB
}

// NewPostRepository initializes and returns a new PostRepository object
func NewPostRepository(db *gorm.DB) *PostRepository {
    return &PostRepository{db}
}

// CreatePost creates a new post
func (pr *PostRepository) CreatePost(post *models.Post) error {
    return pr.db.Create(post).Error
}

// DeletePost deletes a post by ID
func (pr *PostRepository) DeletePost(id uint) error {
    return pr.db.Delete(&models.Post{}, id).Error
}

// GetAllPosts retrieves all posts
func (pr *PostRepository) GetAllPosts() ([]models.Post, error) {
    var posts []models.Post
    if err := pr.db.Find(&posts).Error; err != nil {
        return nil, err
    }
    return posts, nil
}

// GetPostByID retrieves a post by ID
func (pr *PostRepository) GetPostByID(id uint) (*models.Post, error) {
    var post models.Post
    if err := pr.db.First(&post, id).Error; err != nil {
        return nil, err
    }
    return &post, nil
}

// UpdatePost updates a post
func (pr *PostRepository) UpdatePost(post *models.Post) error {
    return pr.db.Save(post).Error
}
