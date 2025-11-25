package createcategory

import (
	"errors"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/repositories/categories"
)

type Service struct {
	repository CategoryRepository
}

func NewService(repository CategoryRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(req *domain.CreateCategoryRequest) (*domain.Category, error) {
	found, err := s.repository.GetCategoryByCode(req.Code)
	if err != nil {
		if !errors.Is(err, categories.ErrCategoryNotFound) {
			return nil, err
		}
	}

	if found != nil {
		return nil, ErrCategoryAlreadyExists
	}

	ctg := &domain.Category{
		Code: req.Code,
		Name: req.Name,
	}

	if err = s.repository.Create(ctg); err != nil {
		return nil, err
	}

	return ctg, nil
}
