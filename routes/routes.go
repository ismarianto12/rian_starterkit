package routes

import (
	"raenRestApi/internal/controller"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	barangController := controller.NewBarangController(db)

	api := r.Group("/api")
	{
		api.POST("/jenisbarang", barangController.CreateJenisBarang)
	}
}
