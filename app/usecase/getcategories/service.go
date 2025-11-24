package getcategories

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type GetCategories struct {
	repository CategoryRepository
}

func NewGetCatalog(r CategoryRepository) *GetCategories {
	return &GetCategories{
		repository: r,
	}
}

func (s *GetCategories) Get() ([]domain.Category, error) {
	cat, err := s.repository.Get()
	if err != nil {
		return nil, err
	}

	return cat, err
}
