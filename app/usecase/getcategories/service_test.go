package getcategories

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/app/usecase/getcategories/mock"
	"github.com/stretchr/testify/assert"
)

type mockOptions struct {
	repository CategoryRepository
}

func TestService_Get(t *testing.T) {
	type want struct {
		res []domain.Category
		err error
	}
	tests := []struct {
		name  string
		want  want
		mocks func(t *testing.T, want want) *mockOptions
	}{
		{
			name: "Error when getting categories",
			want: want{
				res: nil,
				err: errors.New("database error"),
			},
			mocks: func(t *testing.T, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCategoryRepository(ctrl)

				repo.EXPECT().Get().Return(nil, errors.New("database error"))

				return &mockOptions{
					repository: repo,
				}
			},
		},
		{
			name: "Success when getting categories with single category",
			want: want{
				res: []domain.Category{
					{
						ID:   1,
						Code: "CATEGORY_CODE_1",
						Name: "Category 1",
					},
				},
				err: nil,
			},
			mocks: func(t *testing.T, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCategoryRepository(ctrl)

				repo.EXPECT().Get().Return([]domain.Category{
					{
						ID:   1,
						Code: "CATEGORY_CODE_1",
						Name: "Category 1",
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
			opts := test.mocks(t, test.want)

			s := NewGetCatalog(opts.repository)

			res, err := s.Get()

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
