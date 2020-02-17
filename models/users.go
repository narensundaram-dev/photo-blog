package models

import (
	"github.com/jinzhu/gorm"
	// tell gorm that you are going to use sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User Model
type User struct {
	gorm.Model
	Name     string `goorm:"column:name;not null"`
	Username string `gorm:"unique;not null"`
	Password string `goorm:"column:password;not null"`
	Token    string `goorm:"column:token"`
}
