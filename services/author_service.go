package services

import (
	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	responsedto "github.com/GesaXB/LibayGoManagement/dto/responseDto"
	"github.com/GesaXB/LibayGoManagement/models"
	"github.com/GesaXB/LibayGoManagement/repositories"
)

type AuthorService interface {
	GetAll() ([]responsedto.AuthorResponse, error)
	GetById(id uint) (responsedto.AuthorResponse, error)
	Create(req requestdto.AuthorRequest) (responsedto.AuthorResponse, error)
	Update(id uint, req requestdto.AuthorRequest) (responsedto.AuthorResponse, error)
}

type authorService struct {
	repo repositories.AuthorRepository
}

func NewAuthorService(r repositories.AuthorRepository) AuthorService {
	return &authorService{
		repo: r,
	}
}

func (s authorService) GetAll() ([]responsedto.AuthorResponse, error) {
	authors, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]responsedto.AuthorResponse, 0, len(authors))
	for _, author := range authors {
		response = append(response, responsedto.AuthorResponse{
			Name: author.Name,
			Bio:  author.Bio,
		})
	}

	return response, nil
}

func (s authorService) GetById(id uint) (responsedto.AuthorResponse, error) {
	author, err := s.repo.FindById(id)
	if err != nil {
		return responsedto.AuthorResponse{}, err
	}

	response := responsedto.AuthorResponse{
		Name: author.Name,
		Bio:  author.Bio,
	}

	return response, nil
}

func (s authorService) Create(req requestdto.AuthorRequest) (responsedto.AuthorResponse, error) {
	newAuthor := models.Author{
		Name: req.Name,
		Bio:  req.Bio,
	}

	err := s.repo.Create(&newAuthor)
	res := responsedto.AuthorResponse{
		Name: newAuthor.Name,
		Bio:  newAuthor.Bio,
	}

	return res, err
}

func (s authorService) Update(id uint, req requestdto.AuthorRequest) (responsedto.AuthorResponse, error) {
	auhtor, err := s.repo.FindById(id)
	if err != nil {
		return responsedto.AuthorResponse{}, err
	}

	auhtor.Name = req.Name
	auhtor.Bio = req.Bio

	res := responsedto.AuthorResponse{
		Name: auhtor.Name,
		Bio:  auhtor.Bio,
	}

	return res, nil
}
