package domain

import (
	"github.com/shopspring/decimal"
	"strconv"
)

func StringToInt(value string, defaults ...int) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		if len(defaults) > 0 {
			return defaults[0]
		}
	}

	return i
}

func StringToUint(value string, defaults ...uint) uint {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		if len(defaults) > 0 {
			return defaults[0]
		}
	}

	return uint(i)
}

func StringToDecimal(value string, defaults ...decimal.Decimal) decimal.Decimal {
	i, err := decimal.NewFromString(value)
	if err != nil {
		if len(defaults) > 0 {
			return defaults[0]
		}
	}

	return i
}
