package migrate

import (
	"github.com/glaubergoncalves/api-estrutura/api/models"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Usuario{})
}
