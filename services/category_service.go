package services

import (
	"github.com/GesaXB/LibayGoManagement/models"
	"github.com/GesaXB/LibayGoManagement/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
}

type categoryRepository struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(r repositories.CategoryRepository) CategoryService {
	return &categoryRepository{r}
}

func (s categoryRepository) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}
