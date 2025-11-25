package domain

import (
	"net/url"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestProduct_TableName(t *testing.T) {
	product := &Product{}
	assert.Equal(t, "products", product.TableName())
}

type mockQueryGetter struct {
	values map[string]string
}

func (m *mockQueryGetter) Get(key string) string {
	return m.values[key]
}

func TestNewGetProductsRequest(t *testing.T) {
	tests := []struct {
		name     string
		query    map[string]string
		expected GetProductsRequest
	}{
		{
			name: "All parameters provided",
			query: map[string]string{
				"price":       "10.50",
				"category_id": "5",
				"limit":       "20",
				"offset":      "10",
			},
			expected: GetProductsRequest{
				Price:      decimal.RequireFromString("10.50"),
				CategoryID: 5,
				Limit:      20,
				Offset:     10,
			},
		},
		{
			name: "Only required parameters",
			query: map[string]string{
				"price":       "15.75",
				"category_id": "3",
			},
			expected: GetProductsRequest{
				Price:      decimal.RequireFromString("15.75"),
				CategoryID: 3,
				Limit:      10,
				Offset:     0,
			},
		},
		{
			name:  "Empty query parameters",
			query: map[string]string{},
			expected: GetProductsRequest{
				Price:      decimal.Zero,
				CategoryID: 0,
				Limit:      10,
				Offset:     0,
			},
		},
		{
			name: "Invalid price uses default",
			query: map[string]string{
				"price":       "invalid",
				"category_id": "1",
				"limit":       "5",
				"offset":      "2",
			},
			expected: GetProductsRequest{
				Price:      decimal.Zero,
				CategoryID: 1,
				Limit:      5,
				Offset:     2,
			},
		},
		{
			name: "Invalid category_id uses default",
			query: map[string]string{
				"price":       "10.50",
				"category_id": "invalid",
				"limit":       "5",
				"offset":      "2",
			},
			expected: GetProductsRequest{
				Price:      decimal.RequireFromString("10.50"),
				CategoryID: 0,
				Limit:      5,
				Offset:     2,
			},
		},
		{
			name: "Invalid limit uses default",
			query: map[string]string{
				"price":       "10.50",
				"category_id": "1",
				"limit":       "invalid",
				"offset":      "2",
			},
			expected: GetProductsRequest{
				Price:      decimal.RequireFromString("10.50"),
				CategoryID: 1,
				Limit:      10,
				Offset:     2,
			},
		},
		{
			name: "Invalid offset uses default",
			query: map[string]string{
				"price":       "10.50",
				"category_id": "1",
				"limit":       "5",
				"offset":      "invalid",
			},
			expected: GetProductsRequest{
				Price:      decimal.RequireFromString("10.50"),
				CategoryID: 1,
				Limit:      5,
				Offset:     0,
			},
		},
		{
			name: "Limit exceeds maximum (100)",
			query: map[string]string{
				"price":       "10.50",
				"category_id": "1",
				"limit":       "150",
				"offset":      "0",
			},
			expected: GetProductsRequest{
				Price:      decimal.RequireFromString("10.50"),
				CategoryID: 1,
				Limit:      100,
				Offset:     0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryGetter := &mockQueryGetter{values: tt.query}
			result := NewGetProductsRequest(queryGetter)

			assert.Equal(t, tt.expected.Price.String(), result.Price.String())
			assert.Equal(t, tt.expected.CategoryID, result.CategoryID)
			assert.Equal(t, tt.expected.Limit, result.Limit)
			assert.Equal(t, tt.expected.Offset, result.Offset)
		})
	}
}

func TestNewGetProductsRequest_WithURLValues(t *testing.T) {
	values := url.Values{}
	values.Set("price", "25.99")
	values.Set("category_id", "7")
	values.Set("limit", "15")
	values.Set("offset", "5")

	queryGetter := &mockQueryGetter{
		values: map[string]string{
			"price":       values.Get("price"),
			"category_id": values.Get("category_id"),
			"limit":       values.Get("limit"),
			"offset":      values.Get("offset"),
		},
	}

	result := NewGetProductsRequest(queryGetter)

	assert.Equal(t, "25.99", result.Price.String())
	assert.Equal(t, uint(7), result.CategoryID)
	assert.Equal(t, 15, result.Limit)
	assert.Equal(t, 5, result.Offset)
}
