package models

import (
	"github.com/jinzhu/gorm"
	// tell gorm that you are going to use sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Photo Model
type Photo struct {
	gorm.Model
	Path string
	// User int64
}
