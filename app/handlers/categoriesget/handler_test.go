package categoriesget

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/app/handlers/categoriesget/mock"
	"github.com/stretchr/testify/assert"
)

type mockOptions struct {
	getter    CategoriesGetter
	responder Responder
}

func TestHandler_HandleGet(t *testing.T) {
	type want struct {
		statusCode int
	}

	tests := []struct {
		name  string
		want  want
		mocks func(t *testing.T) *mockOptions
	}{
		{
			name: "Error when getting categories",
			want: want{
				statusCode: http.StatusInternalServerError,
			},
			mocks: func(t *testing.T) *mockOptions {
				ctrl := gomock.NewController(t)
				getter := mock.NewMockCategoriesGetter(ctrl)
				responder := mock.NewMockResponder(ctrl)

				dbErr := errors.New("database error")

				getter.EXPECT().
					Get().
					Return(nil, dbErr)

				responder.EXPECT().
					Error(gomock.Any(), http.StatusInternalServerError, dbErr.Error()).
					Do(func(w http.ResponseWriter, status int, message string) {
						w.WriteHeader(status)
					})

				responder.EXPECT().
					Ok(gomock.Any(), gomock.Any()).
					Times(0)

				return &mockOptions{
					getter:    getter,
					responder: responder,
				}
			},
		},
		{
			name: "Success when getting categories with multiple categories",
			want: want{
				statusCode: http.StatusOK,
			},
			mocks: func(t *testing.T) *mockOptions {
				ctrl := gomock.NewController(t)
				getter := mock.NewMockCategoriesGetter(ctrl)
				responder := mock.NewMockResponder(ctrl)

				expected := []domain.Category{
					{
						ID:   1,
						Code: "CATEGORY_CODE_1",
						Name: "Category 1",
					},
					{
						ID:   2,
						Code: "CATEGORY_CODE_2",
						Name: "Category 2",
					},
					{
						ID:   3,
						Code: "CATEGORY_CODE_3",
						Name: "Category 3",
					},
				}

				getter.EXPECT().
					Get().
					Return(expected, nil)

				responder.EXPECT().
					Ok(gomock.Any(), expected).
					Do(func(w http.ResponseWriter, data any) {
						w.WriteHeader(http.StatusOK)
					})

				responder.EXPECT().
					Error(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)

				return &mockOptions{
					getter:    getter,
					responder: responder,
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := test.mocks(t)

			handler := NewHandler(opts.getter, opts.responder)

			req := httptest.NewRequest(http.MethodGet, "/categories", nil)

			recorder := httptest.NewRecorder()

			handler.HandleGet(recorder, req)

			assert.Equal(t, test.want.statusCode, recorder.Code, "Expected status code to match")
		})
	}
}
