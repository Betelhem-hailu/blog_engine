package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	FullName string
	Email     string `gorm:"unique"`
	Password  string
	Posts []Post
}
