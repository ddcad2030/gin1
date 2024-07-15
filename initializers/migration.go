package initializers

import (
	"example/gin-crud-1/models"
)

func migration() {
	DB.Debug().AutoMigrate(&models.Post{})
}
