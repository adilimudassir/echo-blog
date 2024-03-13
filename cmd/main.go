package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/adilimudassir/echo-blog/internal/controllers"
	"github.com/adilimudassir/echo-blog/internal/db"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Database
	dbInstance, err := db.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize Repository for posts and comments
	postRepo := db.NewPostRepository(dbInstance)
	commentRepo := db.NewCommentRepository(dbInstance)

	// Initialize Controllers
	postController := controllers.NewPostController(postRepo)
	commentController := controllers.NewCommentController(commentRepo, postRepo)

	// Routes
	registerPostRoutes(e, postController)
	registerCommentRoutes(e, commentController)

	// Start server
	log.Fatal(e.Start(":8082"))
}

func registerPostRoutes(e *echo.Echo, pc *controllers.PostController) {
	e.POST("/posts", pc.CreatePost)
	e.GET("/posts/:id", pc.GetPostByID)
	e.GET("/posts", pc.GetAllPosts)
	e.PUT("/posts/:id", pc.UpdatePost)
	e.DELETE("/posts/:id", pc.DeletePost)
}

func registerCommentRoutes(e *echo.Echo, cc *controllers.CommentController) {
	e.POST("/posts/:id/comments", cc.CreateComment)
	e.GET("/posts/:id/comments", cc.GetAllCommentsByPostID)
	e.GET("/comments/:id", cc.GetCommentByID)
	e.PUT("/comments/:id", cc.UpdateComment)
	e.DELETE("/comments/:id", cc.DeleteComment)
}
