package controller

import (
	"net/http"

	"raenRestApi/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BarangController struct {
	DB *gorm.DB
}

func NewBarangController(db *gorm.DB) *BarangController {
	return &BarangController{DB: db}
}

func (BarangController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from BarangController"})
}

func (bc *BarangController) CreateJenisBarang(c *gin.Context) {
	var jenisBarang models.JenisBarang
	if err := c.ShouldBindJSON(&jenisBarang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.DB.Create(&jenisBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, jenisBarang)
}
