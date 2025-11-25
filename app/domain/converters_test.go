package domain

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		defaults []int
		expected int
	}{
		{
			name:     "Valid integer",
			value:    "42",
			defaults: []int{},
			expected: 42,
		},
		{
			name:     "Invalid string with default",
			value:    "invalid",
			defaults: []int{10},
			expected: 10,
		},
		{
			name:     "Invalid string without default",
			value:    "invalid",
			defaults: []int{},
			expected: 0,
		},
		{
			name:     "Empty string with default",
			value:    "",
			defaults: []int{5},
			expected: 5,
		},
		{
			name:     "Empty string without default",
			value:    "",
			defaults: []int{},
			expected: 0,
		},
		{
			name:     "Negative integer",
			value:    "-10",
			defaults: []int{},
			expected: -10,
		},
		{
			name:     "Zero",
			value:    "0",
			defaults: []int{},
			expected: 0,
		},
		{
			name:     "Large integer",
			value:    "999999",
			defaults: []int{},
			expected: 999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToInt(tt.value, tt.defaults...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStringToUint(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		defaults []uint
		expected uint
	}{
		{
			name:     "Valid unsigned integer",
			value:    "42",
			defaults: []uint{},
			expected: 42,
		},
		{
			name:     "Invalid string with default",
			value:    "invalid",
			defaults: []uint{10},
			expected: 10,
		},
		{
			name:     "Invalid string without default",
			value:    "invalid",
			defaults: []uint{},
			expected: 0,
		},
		{
			name:     "Empty string with default",
			value:    "",
			defaults: []uint{5},
			expected: 5,
		},
		{
			name:     "Empty string without default",
			value:    "",
			defaults: []uint{},
			expected: 0,
		},
		{
			name:     "Zero",
			value:    "0",
			defaults: []uint{},
			expected: 0,
		},
		{
			name:     "Large unsigned integer",
			value:    "999999",
			defaults: []uint{},
			expected: 999999,
		},
		{
			name:     "Negative number (should parse as 0 or default)",
			value:    "-10",
			defaults: []uint{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToUint(tt.value, tt.defaults...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStringToDecimal(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		defaults []decimal.Decimal
		expected string
	}{
		{
			name:     "Valid decimal",
			value:    "10.50",
			defaults: []decimal.Decimal{},
			expected: "10.50",
		},
		{
			name:     "Invalid string with default",
			value:    "invalid",
			defaults: []decimal.Decimal{decimal.NewFromInt(5)},
			expected: "5",
		},
		{
			name:     "Invalid string without default",
			value:    "invalid",
			defaults: []decimal.Decimal{},
			expected: "0",
		},
		{
			name:     "Empty string with default",
			value:    "",
			defaults: []decimal.Decimal{decimal.NewFromFloat(3.14)},
			expected: "3.14",
		},
		{
			name:     "Empty string without default",
			value:    "",
			defaults: []decimal.Decimal{},
			expected: "0",
		},
		{
			name:     "Zero",
			value:    "0",
			defaults: []decimal.Decimal{},
			expected: "0",
		},
		{
			name:     "Negative decimal",
			value:    "-10.50",
			defaults: []decimal.Decimal{},
			expected: "-10.50",
		},
		{
			name:     "Large decimal",
			value:    "999999.99",
			defaults: []decimal.Decimal{},
			expected: "999999.99",
		},
		{
			name:     "Decimal with many decimal places",
			value:    "10.123456789",
			defaults: []decimal.Decimal{},
			expected: "10.123456789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToDecimal(tt.value, tt.defaults...)
			expectedDecimal := decimal.RequireFromString(tt.expected)
			assert.True(t, result.Equal(expectedDecimal), "Expected %s but got %s", tt.expected, result.String())
		})
	}
}
