package getcatalogbycode

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/app/handlers/getcatalogbycode/mock"
	"github.com/mytheresa/go-hiring-challenge/app/usecase/catalogbycode"
	"github.com/stretchr/testify/assert"
)

type mockOptions struct {
	getter    GetCatalogByCode
	responder Responder
}

func TestHandler_HandleGet(t *testing.T) {
	type args struct {
		code string
	}
	type want struct {
		statusCode int
	}

	tests := []struct {
		name  string
		args  args
		want  want
		mocks func(t *testing.T, args args) *mockOptions
	}{
		{
			name: "Error when getting product by code",
			args: args{
				code: "TEST_CODE",
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
			mocks: func(t *testing.T, args args) *mockOptions {
				ctrl := gomock.NewController(t)
				getter := mock.NewMockGetCatalogByCode(ctrl)
				responder := mock.NewMockResponder(ctrl)

				dbErr := errors.New("database error")

				getter.EXPECT().
					GetByCode(args.code).
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
			name: "Error when getting product by code not found",
			args: args{
				code: "TEST_CODE_NOT_FOUND",
			},
			want: want{
				statusCode: http.StatusNotFound,
			},
			mocks: func(t *testing.T, args args) *mockOptions {
				ctrl := gomock.NewController(t)
				getter := mock.NewMockGetCatalogByCode(ctrl)
				responder := mock.NewMockResponder(ctrl)

				getter.EXPECT().
					GetByCode(args.code).
					Return(nil, catalogbycode.ErrProductNotFound)

				responder.EXPECT().
					Error(gomock.Any(), http.StatusNotFound, catalogbycode.ErrProductNotFound.Error()).
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
			name: "Success when getting product by code with variants",
			args: args{
				code: "TEST_CODE_WITH_VARIANTS",
			},
			want: want{
				statusCode: http.StatusOK,
			},
			mocks: func(t *testing.T, args args) *mockOptions {
				ctrl := gomock.NewController(t)
				getter := mock.NewMockGetCatalogByCode(ctrl)
				responder := mock.NewMockResponder(ctrl)

				expected := &domain.ProductResponse{
					Code:  args.code,
					Price: 10.0,
					Category: domain.Category{
						ID:   123,
						Code: "CATEGORY_CODE",
						Name: "Test",
					},
					Variants: []domain.VariantResponse{
						{
							ID:        1,
							ProductID: 1,
							Name:      "Variant 1",
							SKU:       "VAR-001",
							Price:     10.0,
						},
						{
							ID:        2,
							ProductID: 1,
							Name:      "Variant 2",
							SKU:       "VAR-002",
							Price:     15.0,
						},
					},
				}

				getter.EXPECT().
					GetByCode(args.code).
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
			opts := test.mocks(t, test.args)

			handler := NewCatalogHandler(opts.getter, opts.responder)

			req := httptest.NewRequest(http.MethodGet, "/catalog/"+test.args.code, nil)
			req.SetPathValue("code", test.args.code)

			recorder := httptest.NewRecorder()

			handler.HandleGet(recorder, req)

			assert.Equal(t, test.want.statusCode, recorder.Code, "Expected status code to match")
		})
	}
}
