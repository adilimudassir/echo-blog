package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/adilimudassir/echo-blog/internal/models"
)

var registeredModels = []interface{}{
    &models.Post{},
    &models.Comment{},
}

// NewDB initializes and returns a new DB object with auto migration for all registered models
func NewDB() (*gorm.DB, error) {
    // Open a new SQLite database connection
    // Replace "test.db" with the path to your SQLite database file
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // AutoMigrate the registered model schemas
    if err := db.AutoMigrate(registeredModels...); err != nil {
        return nil, err
    }

    // Return the *gorm.DB object
    return db, nil
}
