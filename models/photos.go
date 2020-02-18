package models

import (
	"github.com/jinzhu/gorm"
	// tell gorm that you are going to use sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Photo Model
type Photo struct {
	gorm.Model
	Path         string `goorm:"column:path;not null"`
	RelativePath string `goorm:"column:relative_path;not null"`
	UserID       uint   `goorm:"column:user_id;not null"`
	Caption      string `goorm:"column:caption"`
}
