package repository

import (
	"github.com/Mukam21/REST-API_online-marketplace/pkg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error

	GetByLogin(login string) (*models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user *models.User) error {

	return r.DB.Create(user).Error
}

func (r *userRepository) GetByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("login = ?", login).First(&user).Error

	return &user, err
}
