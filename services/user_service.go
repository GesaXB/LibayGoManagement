package services

import (
	"github.com/GesaXB/LibayGoManagement/models"
	"github.com/GesaXB/LibayGoManagement/repositories"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) UserService {
	return &userService{*r}
}

func (s *userService) GetAll() ([]models.User, error) {
	user, err := s.repo.FindAll()
	return user, err
}

func (s *userService) GetById(id uint) (models.User, error) {
	user, err := s.repo.FindById(id)
	return user, err
}
