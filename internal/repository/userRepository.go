package repository

import (
	"raenRestApi/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
