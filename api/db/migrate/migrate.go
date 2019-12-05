package migrate

import (
	"github.com/glaubergoncalves/go-api-jwt-token/api/models"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Usuario{})
}
