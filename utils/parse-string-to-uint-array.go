package utils

import (
	"strconv"
	"strings"
)

func StringToUintArray(input string) []uint {
	if input == "" {
		return nil
	}

	parts := strings.Split(input, ",")
	result := make([]uint, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.ParseUint(part, 10, 32)
		if err == nil {
			result = append(result, uint(num))
		}
	}

	return result
}
