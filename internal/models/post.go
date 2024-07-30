package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
    gorm.Model
	Id         int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
	Created Time
	Updated Time
}


type Time struct {
	T time.Time
}

