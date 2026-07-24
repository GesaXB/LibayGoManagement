package services

import (
	responsedto "github.com/GesaXB/LibayGoManagement/dto/responseDto"
	"github.com/GesaXB/LibayGoManagement/repositories"
)

type UserService interface {
	GetAll() ([]responsedto.UserResponse, error)
	GetById(id string) (responsedto.UserResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetAll() ([]responsedto.UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]responsedto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, responsedto.UserResponse{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return responses, nil
}

func (s *userService) GetById(id string) (responsedto.UserResponse, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return responsedto.UserResponse{}, err
	}

	response := responsedto.UserResponse{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil
}
