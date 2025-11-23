package domain

import (
	"strconv"
)

func StringToInt64(value string, defaults ...int64) int64 {
	i, err := strconv.Atoi(value)
	if err != nil {
		if len(defaults) > 0 {
			return defaults[0]
		}
	}

	return int64(i)
}
