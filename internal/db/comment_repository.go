// internal/db/comment_repository.go

package db

import (
    "github.com/adilimudassir/echo-blog/internal/models"
    "gorm.io/gorm"
)

// CommentRepository represents the repository for interacting with comments
type CommentRepository struct {
    db *gorm.DB
}

// NewCommentRepository initializes and returns a new CommentRepository object
func NewCommentRepository(db *gorm.DB) *CommentRepository {
    return &CommentRepository{db}
}

// CreateComment creates a new comment
func (cr *CommentRepository) CreateComment(comment *models.Comment) error {
    return cr.db.Create(comment).Error
}

// GetCommentByID retrieves a comment by its ID
func (cr *CommentRepository) GetCommentByID(id uint) (*models.Comment, error) {
    var comment models.Comment
    if err := cr.db.First(&comment, id).Error; err != nil {
        return nil, err
    }
    return &comment, nil
}

// GetAllCommentsByPostID retrieves all comments for a given post ID
func (cr *CommentRepository) GetAllCommentsByPostID(postID uint) ([]models.Comment, error) {
    var comments []models.Comment
    if err := cr.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
        return nil, err
    }
    return comments, nil
}

// UpdateComment updates an existing comment
func (cr *CommentRepository) UpdateComment(comment *models.Comment) error {
    return cr.db.Save(comment).Error
}

// DeleteComment deletes a comment by its ID
func (cr *CommentRepository) DeleteComment(id uint) error {
    return cr.db.Delete(&models.Comment{}, id).Error
}

// PreloadPost preloads the related Post for a given comment
func (cr *CommentRepository) PreloadPost(comment *models.Comment) error {
    return cr.db.Preload("Post").Find(comment).Error
}