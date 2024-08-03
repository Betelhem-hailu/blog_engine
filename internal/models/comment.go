package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
    gorm.Model
	Id         int    `json:"id"`
    Name   string `json:"name"`
    Message string `json:"message"`
}
