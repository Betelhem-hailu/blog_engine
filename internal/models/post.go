package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
    gorm.Model
	Id        uint `gorm:"primaryKey"`
    Title   string 
    Content string 
	UserID  uint   // Foreign key for User
	User    User 
}

