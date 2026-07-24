package repositories

import (
	"github.com/GesaXB/LibayGoManagement/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindById(id uint) (models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r categoryRepository) FindById(id uint) (models.Category, error) {
	var category models.Category

	err := r.db.First(&category, id).Error

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r categoryRepository) Create(category *models.Category) error {
	err := r.db.Create(category).Error
	return err
}

func (r categoryRepository) Update(category *models.Category) error {
	err := r.db.Save(category).Error
	return err
}
