package services

import (
	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	responsedto "github.com/GesaXB/LibayGoManagement/dto/responseDto"
	"github.com/GesaXB/LibayGoManagement/models"
	"github.com/GesaXB/LibayGoManagement/repositories"
)

type BookService interface {
	GetAllBooks() ([]responsedto.BookResponse, error)
	GetBookById(id string) (responsedto.BookResponse, error)
	CreateBook(book requestdto.BookRequest) (responsedto.BookResponse, error)
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]responsedto.BookResponse, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	responses := make([]responsedto.BookResponse, len(books))
	for i, book := range books {
		responses[i] = responsedto.BookResponse{
			ID:          book.Id,
			Title:       book.Title,
			Isbnd:       book.Isbnd,
			Description: book.Description,
			Stock:       book.Stock,
			Author: responsedto.AuthorResponse{
				Id:   book.AuthorId,
				Name: book.Author.Name,
				Bio:  book.Author.Bio,
			},
			Category: responsedto.CategoryRespone{
				Id:   book.CategoryId,
				Name: book.Category.Name},
		}
	}
	return responses, nil
}

func (s *bookService) GetBookById(id string) (responsedto.BookResponse, error) {
	book, err := s.repo.GetById(id)
	if err != nil {
		return responsedto.BookResponse{}, err
	}
	return responsedto.BookResponse{
		ID:          book.Id,
		Title:       book.Title,
		Isbnd:       book.Isbnd,
		Description: book.Description,
		Stock:       book.Stock,
		Author: responsedto.AuthorResponse{
			Id:   book.AuthorId,
			Name: book.Author.Name,
			Bio:  book.Author.Bio},
		Category: responsedto.CategoryRespone{
			Id:   book.CategoryId,
			Name: book.Category.Name,
		},
	}, nil
}

func (s *bookService) CreateBook(book requestdto.BookRequest) (responsedto.BookResponse, error) {
	bookModel := models.Book{
		Title:       book.Title,
		Isbnd:       book.Isbnd,
		Description: book.Description,
		Stock:       book.Stock,
		AuthorId:    book.AuthorId,
		CategoryId:  book.CategoryId,
	}
	err := s.repo.Create(&bookModel)
	if err != nil {
		return responsedto.BookResponse{}, err
	}

	createdBook, err := s.repo.GetById(bookModel.Id)
	if err != nil {
		return responsedto.BookResponse{}, err
	}

	response := responsedto.BookResponse{
		ID:          createdBook.Id,
		Title:       createdBook.Title,
		Isbnd:       createdBook.Isbnd,
		Description: createdBook.Description,
		Stock:       createdBook.Stock,
		Author: responsedto.AuthorResponse{
			Id:   createdBook.AuthorId,
			Name: createdBook.Author.Name,
			Bio:  createdBook.Author.Bio,
		},
		Category: responsedto.CategoryRespone{
			Id:   createdBook.CategoryId,
			Name: createdBook.Category.Name,
		},
	}
	return response, nil
}
