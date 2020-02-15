package db

import (
	"github.com/jinzhu/gorm"
	// Tell gorm to use sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"photo-blog/models"
)

// User Model
type User models.User

// Photo Model
type Photo models.Photo

// DB Global variable
var DB *gorm.DB

func migrate(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(
		&User{},
		&Photo{},
	)
}

// OpenDBConnection opens db connection
func OpenDBConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "photo-blog.db")
	if err != nil {
		panic("failed to connect database")
	}

	migrate(db)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// db.LogMode(true)

	DB = db
	return db
}

// Get returns sqlite db drive
func Get() *gorm.DB {
	if DB == nil {
		panic("DB is not initialized!")
	} else {
		return DB
	}
}
