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

func (r *Repository) Create(ctg *domain.Category) error {
	if err := r.db.Create(&ctg).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get() ([]domain.Category, error) {
	var cat []domain.Category

	if err := r.db.Model(&domain.Category{}).Find(&cat).Error; err != nil {
		return nil, err
	}

	return cat, nil
}

func (r *Repository) GetCategoryByCode(code string) (*domain.Category, error) {
	var cat domain.Category

	if err := r.db.Model(&cat).
		First(&cat, "code = ?", code).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return &cat, nil
}
