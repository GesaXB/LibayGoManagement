package services

import (
	"errors"

	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	"github.com/GesaXB/LibayGoManagement/models"
	"github.com/GesaXB/LibayGoManagement/repositories"
	"github.com/GesaXB/LibayGoManagement/utils"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req requestdto.RegisterRequest) (string, error)
	Login(req requestdto.LoginRequest) (string, *models.User, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(req requestdto.RegisterRequest) (string, error) {

	user, err := s.repo.FindByEmail(req.Email)

	if err == nil && user != nil {
		return "", errors.New("email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	newUser := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.repo.Create(&newUser); err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(newUser.Id, newUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Login(req requestdto.LoginRequest) (string, *models.User, error) {

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	token, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
