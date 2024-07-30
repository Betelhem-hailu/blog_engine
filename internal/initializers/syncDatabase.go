package initializers

import (
	"fmt"
	"go-gin-blog/internal/models"
	"log"
)

func SyncDatabase() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("Error migrating database: %v\n", err)
    } else {
        fmt.Println("Database migrated successfully!")
    }
}