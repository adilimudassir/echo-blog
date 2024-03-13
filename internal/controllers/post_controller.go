package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
	"github.com/adilimudassir/echo-blog/internal/db"
	"github.com/adilimudassir/echo-blog/internal/models"
	"github.com/adilimudassir/echo-blog/internal/helpers"
)

// PostController represents the controller for handling posts
type PostController struct {
	postRepo *db.PostRepository
	validate *validator.Validate
}

// NewPostController initializes and returns a new PostController object
func NewPostController(postRepo *db.PostRepository) *PostController {
	return &PostController{
		postRepo: postRepo,
		validate: validator.New(),
	}
}

// CreatePost creates a new post
func (pc *PostController) CreatePost(c echo.Context) error {
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return err
	}

	// Validate the post
	if err := pc.validate.Struct(post); err != nil {
        return c.JSON(http.StatusBadRequest, helpers.FormatValidationErrors(err))
    }

	if err := pc.postRepo.CreatePost(post); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, post)
}

// GetPostByID retrieves a post by ID
func (pc *PostController) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	post, err := pc.postRepo.GetPostByID(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, post)
}

// GetAllPosts retrieves all posts
func (pc *PostController) GetAllPosts(c echo.Context) error {
	posts, err := pc.postRepo.GetAllPosts()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, posts)
}

// UpdatePost updates a post
func (pc *PostController) UpdatePost(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return err
    }

    post := new(models.Post)
    if err := c.Bind(post); err != nil {
        return err
    }
    post.ID = uint(id)

    // Validate the post
    if err := pc.validate.Struct(post); err != nil {
        return c.JSON(http.StatusBadRequest, helpers.FormatValidationErrors(err))
    }

    if err := pc.postRepo.UpdatePost(post); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post by ID
func (pc *PostController) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := pc.postRepo.DeletePost(uint(id)); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
