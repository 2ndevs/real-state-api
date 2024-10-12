package libs

import "strconv"

func ValidateAndConvertCoordinate(value string, min, max float64) *float64 {
	if value == "" {
		return nil
	}

	coordinate, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil
	}

	if coordinate < min || coordinate > max {
		return nil
	}

	return &coordinate
}
