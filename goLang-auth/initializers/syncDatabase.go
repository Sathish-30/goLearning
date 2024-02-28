package initializers

import (
	"github.com/sathish-30/golang-auth/internal/database"
	"github.com/sathish-30/golang-auth/internal/models"
)

func SyncDatabase() {
	database.DB.AutoMigrate(&models.User{})
}
