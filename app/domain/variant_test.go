package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariant_TableName(t *testing.T) {
	variant := &Variant{}
	assert.Equal(t, "product_variants", variant.TableName())
}
