package services

import (
	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	responsedto "github.com/GesaXB/LibayGoManagement/dto/responseDto"
	"github.com/GesaXB/LibayGoManagement/models"
	"github.com/GesaXB/LibayGoManagement/repositories"
)

type CategoryService interface {
	GetAll() ([]responsedto.CategoryRespone, error)
	GetById(id string) (responsedto.CategoryRespone, error)
	Create(req requestdto.CategoryRequest) (responsedto.CategoryRespone, error)
	Update(id string, req requestdto.UpdateCategoryRequest) (responsedto.CategoryRespone, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(r repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repo: r,
	}
}

func (s categoryService) GetAll() ([]responsedto.CategoryRespone, error) {
	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]responsedto.CategoryRespone, 0, len(categories))
	for _, category := range categories {
		responses = append(responses, responsedto.CategoryRespone{
			Id:   category.Id,
			Name: category.Name,
		})
	}
	return responses, nil
}

func (s categoryService) GetById(id string) (responsedto.CategoryRespone, error) {
	category, err := s.repo.FindById(id)
	if err != nil {
		return responsedto.CategoryRespone{}, err
	}

	response := responsedto.CategoryRespone{
		Id:   category.Id,
		Name: category.Name,
	}

	return response, nil
}

func (s categoryService) Create(req requestdto.CategoryRequest) (responsedto.CategoryRespone, error) {

	newCategory := models.Category{
		Name: req.Name,
	}

	err := s.repo.Create(&newCategory)
	res := responsedto.CategoryRespone{
		Id:   newCategory.Id,
		Name: newCategory.Name,
	}

	return res, err
}

func (s categoryService) Update(id string, req requestdto.UpdateCategoryRequest) (responsedto.CategoryRespone, error) {
	category, err := s.repo.FindById(id)
	if err != nil {
		return responsedto.CategoryRespone{}, err
	}

	category.Name = req.Name
	res := responsedto.CategoryRespone{
		Id:   category.Id,
		Name: category.Name,
	}
	return res, nil
}
