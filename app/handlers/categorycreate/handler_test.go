package categorycreate

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/app/handlers/categorycreate/mock"
	"github.com/stretchr/testify/assert"
)

type mockOptions struct {
	creator   CategoriesCreator
	responder Responder
}

func TestHandler_HandlePost(t *testing.T) {
	type want struct {
		statusCode int
	}

	tests := []struct {
		name  string
		body  string
		want  want
		mocks func(t *testing.T, body string) *mockOptions
	}{
		{
			name: "Error when JSON is invalid",
			body: `{"code": "TEST_CODE", "name": invalid}`,
			want: want{
				statusCode: http.StatusBadRequest,
			},
			mocks: func(t *testing.T, body string) *mockOptions {
				ctrl := gomock.NewController(t)
				creator := mock.NewMockCategoriesCreator(ctrl)
				responder := mock.NewMockResponder(ctrl)

				responder.EXPECT().
					Error(gomock.Any(), http.StatusBadRequest, "invalid JSON").
					Do(func(w http.ResponseWriter, status int, message string) {
						w.WriteHeader(status)
					})

				responder.EXPECT().
					Ok(gomock.Any(), gomock.Any()).
					Times(0)

				creator.EXPECT().
					Create(gomock.Any()).
					Times(0)

				return &mockOptions{
					creator:   creator,
					responder: responder,
				}
			},
		},
		{
			name: "Error when validation fails - missing code",
			body: `{"name": "Test Category"}`,
			want: want{
				statusCode: http.StatusBadRequest,
			},
			mocks: func(t *testing.T, body string) *mockOptions {
				ctrl := gomock.NewController(t)
				creator := mock.NewMockCategoriesCreator(ctrl)
				responder := mock.NewMockResponder(ctrl)

				responder.EXPECT().
					Error(gomock.Any(), http.StatusBadRequest, gomock.Any()).
					Do(func(w http.ResponseWriter, status int, message string) {
						w.WriteHeader(status)
						assert.Contains(t, message, "code")
					})

				responder.EXPECT().
					Ok(gomock.Any(), gomock.Any()).
					Times(0)

				creator.EXPECT().
					Create(gomock.Any()).
					Times(0)

				return &mockOptions{
					creator:   creator,
					responder: responder,
				}
			},
		},
		{
			name: "Error when validation fails - missing name",
			body: `{"code": "TEST_CODE"}`,
			want: want{
				statusCode: http.StatusBadRequest,
			},
			mocks: func(t *testing.T, body string) *mockOptions {
				ctrl := gomock.NewController(t)
				creator := mock.NewMockCategoriesCreator(ctrl)
				responder := mock.NewMockResponder(ctrl)

				responder.EXPECT().
					Error(gomock.Any(), http.StatusBadRequest, gomock.Any()).
					Do(func(w http.ResponseWriter, status int, message string) {
						w.WriteHeader(status)
						assert.Contains(t, message, "name")
					})

				responder.EXPECT().
					Ok(gomock.Any(), gomock.Any()).
					Times(0)

				creator.EXPECT().
					Create(gomock.Any()).
					Times(0)

				return &mockOptions{
					creator:   creator,
					responder: responder,
				}
			},
		},
		{
			name: "Error when validation fails - name too long",
			body: `{"code": "TEST_CODE", "name": "` + strings.Repeat("a", 201) + `"}`,
			want: want{
				statusCode: http.StatusBadRequest,
			},
			mocks: func(t *testing.T, body string) *mockOptions {
				ctrl := gomock.NewController(t)
				creator := mock.NewMockCategoriesCreator(ctrl)
				responder := mock.NewMockResponder(ctrl)

				responder.EXPECT().
					Error(gomock.Any(), http.StatusBadRequest, gomock.Any()).
					Do(func(w http.ResponseWriter, status int, message string) {
						w.WriteHeader(status)
					})

				responder.EXPECT().
					Ok(gomock.Any(), gomock.Any()).
					Times(0)

				creator.EXPECT().
					Create(gomock.Any()).
					Times(0)

				return &mockOptions{
					creator:   creator,
					responder: responder,
				}
			},
		},
		{
			name: "Error when creating category",
			body: `{"code": "TEST_CODE", "name": "Test Category"}`,
			want: want{
				statusCode: http.StatusInternalServerError,
			},
			mocks: func(t *testing.T, body string) *mockOptions {
				ctrl := gomock.NewController(t)
				creator := mock.NewMockCategoriesCreator(ctrl)
				responder := mock.NewMockResponder(ctrl)

				dbErr := errors.New("database error")

				creator.EXPECT().
					Create(gomock.Any()).
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
					creator:   creator,
					responder: responder,
				}
			},
		},
		{
			name: "Success when creating category",
			body: `{"code": "TEST_CODE", "name": "Test Category"}`,
			want: want{
				statusCode: http.StatusOK,
			},
			mocks: func(t *testing.T, body string) *mockOptions {
				ctrl := gomock.NewController(t)
				creator := mock.NewMockCategoriesCreator(ctrl)
				responder := mock.NewMockResponder(ctrl)

				var req domain.CreateCategoryRequest
				json.Unmarshal([]byte(body), &req)

				expected := &domain.Category{
					Code: req.Code,
					Name: req.Name,
				}

				creator.EXPECT().
					Create(gomock.Any()).
					Do(func(r *domain.CreateCategoryRequest) {
						assert.Equal(t, req.Code, r.Code)
						assert.Equal(t, req.Name, r.Name)
					}).
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
					creator:   creator,
					responder: responder,
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := test.mocks(t, test.body)

			handler := NewHandler(opts.creator, opts.responder)

			req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBufferString(test.body))
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()

			handler.HandlePost(recorder, req)

			assert.Equal(t, test.want.statusCode, recorder.Code, "Expected status code to match")
		})
	}
}
