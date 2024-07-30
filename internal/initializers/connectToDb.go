package initializers

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("/home/b/Documents/personal_Project/blog_engine/internal/sql/blog.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	fmt.Println("Database connected successfully!")

}