package main

import (
	"testing"

	"github.com/ro-mish/GoStruct/pkg/utils"
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
		result := utils.ParseValue(test.input)
		if result != test.expected {
			t.Errorf("ParseValue(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestParseValueEdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"", ""},                 // Empty string
		{"0", 0},                 // Zero value
		{"-123", -123},           // Negative integer
		{"-45.67", -45.67},       // Negative float
		{"3.14159", 3.14159},     // Pi
		{"999999999", 999999999}, // Large integer
		{"abc123", "abc123"},     // Alphanumeric
		{"!@#$", "!@#$"},         // Special characters
		{"0.0", 0.0},             // Zero float
		{"00123", 123},           // Leading zeros
	}

	for _, test := range tests {
		result := utils.ParseValue(test.input)
		if result != test.expected {
			t.Errorf("ParseValue(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestParseValueInvalidNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"1.2.3", "1.2.3"},   // Multiple decimal points
		{"1e999", "1e999"},   // Number too large
		{".123", ".123"},     // Missing leading zero
		{"123.", "123."},     // Trailing decimal
		{"12,345", "12,345"}, // Comma in number
	}

	for _, test := range tests {
		result := utils.ParseValue(test.input)
		if result != test.expected {
			t.Errorf("ParseValue(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
