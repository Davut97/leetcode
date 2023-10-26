package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	// Test cases
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994}, // Edge case: empty string
	}

	// Iterate through the test cases
	for _, test := range tests {
		reversed := RomanToInt(test.input)
		if reversed != test.expected {
			t.Errorf("RomanToInt(%s) = %v; expected %v", test.input, reversed, test.expected)
		}
	}
}
