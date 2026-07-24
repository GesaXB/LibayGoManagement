package repositories

import (
	"github.com/GesaXB/LibayGoManagement/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var listUser []models.User

	err := r.db.Find(&listUser).Error
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (r *userRepository) FindById(id string) (models.User, error) {
	var user models.User

	err := r.db.First(&user, "id = ?", id).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	err := r.db.Create(user).Error
	return err
}
