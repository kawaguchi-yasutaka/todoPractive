package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	IsArchive   bool `gorm:"default: false"`
}
