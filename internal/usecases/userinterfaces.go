package usecases

import (
	"raenRestApi/internal/models"
	"raenRestApi/internal/repository"

	"gorm.io/gorm"
)

type UserUsecase struct {
	userRepo *repository.UserRepository
}

func NewUserUsecase(db *gorm.DB) *UserUsecase {
	return &UserUsecase{
		userRepo: repository.NewUserRepository(db),
	}
}

func (uu *UserUsecase) CreateUser(user *models.User) error {
	return uu.userRepo.CreateUser(user)
}

func (uu *UserUsecase) GetUserByID(id uint) (*models.User, error) {
	return uu.userRepo.GetUserByID(id)
}
