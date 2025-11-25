package createcategory

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/app/usecase/createcategory/mock"
	"github.com/mytheresa/go-hiring-challenge/repositories/categories"
	"github.com/stretchr/testify/assert"
)

type mockOptions struct {
	repository CategoryRepository
}

func TestService_Create(t *testing.T) {
	type args struct {
		req *domain.CreateCategoryRequest
	}
	type want struct {
		res *domain.Category
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(t *testing.T, args args, want want) *mockOptions
	}{
		{
			name: "Error when getting category by code returns database error",
			args: args{
				req: &domain.CreateCategoryRequest{
					Code: "TEST_CODE",
					Name: "Test Category",
				},
			},
			want: want{
				res: nil,
				err: errors.New("database error"),
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCategoryRepository(ctrl)

				repo.EXPECT().
					GetCategoryByCode(args.req.Code).
					Return(nil, errors.New("database error"))

				return &mockOptions{
					repository: repo,
				}
			},
		},
		{
			name: "Error when category already exists",
			args: args{
				req: &domain.CreateCategoryRequest{
					Code: "EXISTING_CODE",
					Name: "Existing Category",
				},
			},
			want: want{
				res: nil,
				err: ErrCategoryAlreadyExists,
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCategoryRepository(ctrl)

				existingCategory := &domain.Category{
					ID:   1,
					Code: "EXISTING_CODE",
					Name: "Existing Category",
				}

				repo.EXPECT().
					GetCategoryByCode(args.req.Code).
					Return(existingCategory, nil)

				return &mockOptions{
					repository: repo,
				}
			},
		},
		{
			name: "Error when creating category in repository",
			args: args{
				req: &domain.CreateCategoryRequest{
					Code: "NEW_CODE",
					Name: "New Category",
				},
			},
			want: want{
				res: nil,
				err: errors.New("create error"),
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCategoryRepository(ctrl)

				repo.EXPECT().
					GetCategoryByCode(args.req.Code).
					Return(nil, categories.ErrCategoryNotFound)

				repo.EXPECT().
					Create(gomock.Any()).
					Return(errors.New("create error"))

				return &mockOptions{
					repository: repo,
				}
			},
		},
		{
			name: "Success when creating new category",
			args: args{
				req: &domain.CreateCategoryRequest{
					Code: "NEW_CODE",
					Name: "New Category",
				},
			},
			want: want{
				res: &domain.Category{
					Code: "NEW_CODE",
					Name: "New Category",
				},
				err: nil,
			},
			mocks: func(t *testing.T, args args, want want) *mockOptions {
				ctrl := gomock.NewController(t)
				repo := mock.NewMockCategoryRepository(ctrl)

				repo.EXPECT().
					GetCategoryByCode(args.req.Code).
					Return(nil, categories.ErrCategoryNotFound)

				repo.EXPECT().
					Create(gomock.Any()).
					Do(func(category *domain.Category) {
						assert.Equal(t, args.req.Code, category.Code)
						assert.Equal(t, args.req.Name, category.Name)
					}).
					Return(nil)

				return &mockOptions{
					repository: repo,
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := test.mocks(t, test.args, test.want)

			s := NewService(opts.repository)

			res, err := s.Create(test.args.req)

			a := assert.New(t)

			if test.want.err != nil {
				a.Error(err)
				if test.want.err == ErrCategoryAlreadyExists {
					a.Equal(test.want.err, err)
				} else {
					a.Contains(err.Error(), test.want.err.Error())
				}
			}

			if test.want.err == nil {
				a.NoError(err)
				a.Equal(test.want.res.Code, res.Code)
				a.Equal(test.want.res.Name, res.Name)
			}
		})
	}
}
