package catalogbycode

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/app/usecase/catalogbycode/mock"
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockOptions struct {
	repository CatalogRepository
}

func TestService_GetCatalogByCode(t *testing.T) {
	type args struct {
		input string
	}
	type want struct {
		res *domain.ProductResponse
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(t *testing.T, args args, want want) *mockOptions
	}{
		{
			name: "Error when getting product by code",
			args: args{
				input: "TEST_CODE",
			},
			want: want{
				res: nil,
				err: errors.New("database error"),
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCatalogRepository(ctrl)

				repo.EXPECT().GetProductByCode("TEST_CODE").Return(nil, errors.New("database error"))

				return &mockOptions{
					repository: repo,
				}
			},
		},
		{
			name: "Error when getting product by code not found",
			args: args{
				input: "TEST_CODE_NOT_FOUND",
			},
			want: want{
				res: nil,
				err: ErrProductNotFound,
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCatalogRepository(ctrl)

				repo.EXPECT().GetProductByCode("TEST_CODE_NOT_FOUND").Return(nil, models.ErrProductNotFound)

				return &mockOptions{
					repository: repo,
				}
			},
		},
		{
			name: "Success when getting product by code",
			args: args{
				input: "TEST_CODE",
			},
			want: want{
				res: &domain.ProductResponse{
					Code:  "TEST_CODE",
					Price: 10.0,
					Category: domain.Category{
						ID:   123,
						Code: "CATEGORY_CODE",
						Name: "Test",
					},
					Variants: []domain.VariantResponse{},
				},
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCatalogRepository(ctrl)

				repo.EXPECT().GetProductByCode("TEST_CODE").Return(&domain.Product{
					Code:  "TEST_CODE",
					Price: decimal.NewFromInt(10),
					Category: domain.Category{
						ID:   123,
						Code: "CATEGORY_CODE",
						Name: "Test",
					},
				}, nil)

				return &mockOptions{
					repository: repo,
				}
			},
		},
		{
			name: "Success when getting product by code without price",
			args: args{
				input: "TEST_CODE",
			},
			want: want{
				res: &domain.ProductResponse{
					Code:  "TEST_CODE",
					Price: 10.0,
					Category: domain.Category{
						ID:   123,
						Code: "CATEGORY_CODE",
						Name: "Test",
					},
					Variants: []domain.VariantResponse{
						domain.VariantResponse{
							Name:  "variant name",
							SKU:   "sku",
							Price: 10,
						},
					},
				},
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCatalogRepository(ctrl)

				repo.EXPECT().GetProductByCode("TEST_CODE").Return(&domain.Product{
					Code:  "TEST_CODE",
					Price: decimal.NewFromInt(10),
					Category: domain.Category{
						ID:   123,
						Code: "CATEGORY_CODE",
						Name: "Test",
					},

					Variants: []domain.Variant{
						domain.Variant{
							Name: "variant name",
							SKU:  "sku",
						},
					},
				}, nil)

				return &mockOptions{
					repository: repo,
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := test.mocks(t, test.args, test.want)

			s := NewGetCatalog(opts.repository)

			res, err := s.GetByCode(test.args.input)

			a := assert.New(t)

			if test.want.err != nil {
				a.Error(err)
				a.Contains(err.Error(), test.want.err.Error())
			}

			if test.want.err == nil {
				a.NoError(err)
				a.Equal(test.want.res, res)
			}
		})
	}
}
