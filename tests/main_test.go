package main

import (
	"testing"
)

func TestParseValue(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"123", 123},
		{"45.67", 45.67},
		{"hello", "hello"},
	}

	for _, test := range tests {
		result := parseValue(test.input)
		if result != test.expected {
			t.Errorf("parseValue(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
