package utils

import "strconv"

func ParseParamToBool(value string) *bool {
	if value == "" {
		return nil
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}

	return &boolValue
}
