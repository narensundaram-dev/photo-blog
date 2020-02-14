package models

import (
	"github.com/jinzhu/gorm"
	// tell gorm that you are going to use sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User Model
type User struct {
	gorm.Model
	Name     string
	Username string
	Password string
}
