package repositories

import (
	"github.com/GesaXB/LibayGoManagement/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	GetById(id string) (models.Book, error)
	Create(*models.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Preload("Author").Preload("Category").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) GetById(id string) (models.Book, error) {
	var book models.Book
	err := r.db.Preload("Author").Preload("Category").First(&book, "id = ?", id).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (r *bookRepository) Create(book *models.Book) error {
	err := r.db.Create(book).Error
	return err
}
