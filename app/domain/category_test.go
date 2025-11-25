package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory_TableName(t *testing.T) {
	category := &Category{}
	assert.Equal(t, "categories", category.TableName())
}

func TestCreateCategoryRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateCategoryRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			request: CreateCategoryRequest{
				Code: "TEST_CODE",
				Name: "Test Category",
			},
			wantErr: false,
		},
		{
			name: "Error when code is missing",
			request: CreateCategoryRequest{
				Code: "",
				Name: "Test Category",
			},
			wantErr: true,
		},
		{
			name: "Error when name is missing",
			request: CreateCategoryRequest{
				Code: "TEST_CODE",
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "Error when name is too long",
			request: CreateCategoryRequest{
				Code: "TEST_CODE",
				Name: string(make([]byte, 201)),
			},
			wantErr: true,
		},
		{
			name: "Valid when name is exactly 200 characters",
			request: CreateCategoryRequest{
				Code: "TEST_CODE",
				Name: string(make([]byte, 200)),
			},
			wantErr: false,
		},
		{
			name: "Valid when name is exactly 1 character",
			request: CreateCategoryRequest{
				Code: "TEST_CODE",
				Name: "A",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
