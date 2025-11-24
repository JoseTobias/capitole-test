package categories

import (
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get() ([]domain.Category, error) {
	var cat []domain.Category

	if err := r.db.Model(&domain.Category{}).Find(&cat).Error; err != nil {
		return nil, err
	}

	return cat, nil
}
