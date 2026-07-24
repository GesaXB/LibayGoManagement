package repositories

import (
	"github.com/GesaXB/LibayGoManagement/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() ([]models.Author, error)
	FindById(id uint) (models.Author, error)
	Create(author *models.Author) error
	Update(author *models.Author) error
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

func (r authorRepository) FindAll() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, err
}

func (r authorRepository) FindById(id uint) (models.Author, error) {
	var author models.Author
	err := r.db.First(id, &author).Error
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}

func (r authorRepository) Create(author *models.Author) error {
	err := r.db.Create(author).Error
	return err
}

func (r authorRepository) Update(author *models.Author) error {
	err := r.db.Save(author).Error
	return err
}
