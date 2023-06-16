package initializers

import (
	"fmt"
	"go-jwt/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	}
}
