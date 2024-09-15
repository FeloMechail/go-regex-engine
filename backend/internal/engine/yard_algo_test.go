package engine

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		// Simple expressions
		{"3 + 4", []string{"3", "4", "+"}},
		{"2 * 3", []string{"2", "3", "*"}},
		{"5 - 2", []string{"5", "2", "-"}},

		// Expressions with parentheses
		{"( 1 + 2 ) * 3", []string{"1", "2", "+", "3", "*"}},
		{"( 4 + 5 ) / ( 7 - 3 )", []string{"4", "5", "+", "7", "3", "-", "/"}},
		{"( 8 / 4 ) + ( 6 * 2 )", []string{"8", "4", "/", "6", "2", "*", "+"}},

		// Expressions with exponentiation
		{"3 + 5 ^ 2", []string{"3", "5", "2", "^", "+"}},
		{"( 1 + 2 ) ^ 3", []string{"1", "2", "+", "3", "^"}},

		// Complex expressions
		{"3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3", []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "^", "^", "/", "+"}},
		{"( 1 + 2 * 3 ) / 4", []string{"1", "2", "3", "*", "+", "4", "/"}},

		// Edge cases
		{"7", []string{"7"}},        // Single number
		{"+ 5", []string{"5", "+"}}, // Operator before operand
		{"7 *", []string{"7", "*"}}, // Operand before operator

	}

	for _, test := range tests {
		result := ParseInput(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ParseInput(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
