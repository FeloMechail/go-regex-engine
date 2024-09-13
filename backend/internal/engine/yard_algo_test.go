package engine

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3+4*2/(1-5)^2^3", "3 4 2 * 1 5 - 2 3 ^ ^ / +"},
		{"3+4*2/(1-5)*2*3", "3 4 2 * 1 5 - / 2 * 3 * +"},
		{"(1+2)*3", "1 2 + 3 *"},
		{"3+4*2/(1-5)", "3 4 2 * 1 5 - / +"},
		{"3+4*2", "3 4 2 * +"},
	}

	for _, test := range tests {
		result := ParseInput(test.input)
		if result != test.expected {
			t.Errorf("ParseInput(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
