package utils

import (
	"strconv"
	"strings"
)

func ParseValue(value string) interface{} {
	var intValue int
	var floatValue float64
	var err error

	// Check for edge cases where the value should remain a string
	if strings.HasPrefix(value, ".") || strings.HasSuffix(value, ".") || strings.Contains(value, ",") {
		return value
	}

	// Try to parse the value as an integer
	if intValue, err = strconv.Atoi(value); err == nil {
		return intValue
	}

	// Try to parse the value as a float
	if floatValue, err = strconv.ParseFloat(value, 64); err == nil {
		return floatValue
	}

	return value
}
