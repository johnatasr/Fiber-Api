package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int64  `json:"rating"`
}
