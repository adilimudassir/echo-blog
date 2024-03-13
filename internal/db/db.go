// internal/db/db.go

package db

import (
	"github.com/adilimudassir/echo-blog/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB represents the interface for interacting with the database
type DB interface {
    PostRepository() PostRepository
}

// NewDB initializes and returns a new DB object
func NewDB() (*gorm.DB, error) {
    // Open a new SQLite database connection
    // Replace "test.db" with the path to your SQLite database file
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // AutoMigrate the model schemas
    if err := db.AutoMigrate(&models.Post{}); err != nil {
        return nil, err
    }

    // Return the *gorm.DB object
    return db, nil
}
