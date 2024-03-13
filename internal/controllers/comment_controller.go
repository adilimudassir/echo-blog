package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adilimudassir/echo-blog/internal/db"
	"github.com/adilimudassir/echo-blog/internal/helpers"
	"github.com/adilimudassir/echo-blog/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CommentController represents the controller for handling comments
type CommentController struct {
	repo     *db.CommentRepository
	postRepo *db.PostRepository
	validate *validator.Validate
}

// NewCommentController initializes and returns a new CommentController object
func NewCommentController(repo *db.CommentRepository, postRepo *db.PostRepository) *CommentController {
	return &CommentController{
		repo:     repo,
		postRepo: postRepo,
		validate: validator.New(),
	}
}

// CreateComment creates a new comment
func (cc *CommentController) CreateComment(c echo.Context) error {
	// Retrieve post ID from route parameters
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Find the corresponding post
	post, err := cc.postRepo.GetPostByID(uint(postID))
	if err != nil {
		return err
	}

	// Bind the comment from the request body
	comment := new(models.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}

	// Validate the comment
	if err := cc.validate.Struct(comment); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FormatValidationErrors(err))
	}

	// Associate the comment with the post
	comment.PostID = post.ID

	// Create the comment
	if err := cc.repo.CreateComment(comment); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, comment)
}

// GetCommentByID retrieves a comment by its ID
func (cc *CommentController) GetCommentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	comment, err := cc.repo.GetCommentByID(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

// GetAllCommentsByPostID retrieves all comments for a given post ID
func (cc *CommentController) GetAllCommentsByPostID(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return err
	}
	comments, err := cc.repo.GetAllCommentsByPostID(uint(postID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, comments)
}

// UpdateComment updates an existing comment
func (cc *CommentController) UpdateComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	comment := new(models.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	comment.ID = uint(id)

	// Validate the comment
	if err := cc.validate.Struct(comment); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FormatValidationErrors(err))
	}

	if err := cc.repo.UpdateComment(comment); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

// DeleteComment deletes a comment by its ID
func (cc *CommentController) DeleteComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if err := cc.repo.DeleteComment(uint(id)); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
