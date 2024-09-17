package engine

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		// Simple concatenation
		{"ab", []string{"a", "b", "&"}},            // "a" followed by "b" -> RPN: a b &
		{"abc", []string{"a", "b", "&", "c", "&"}}, // "a" followed by "b" then "c" -> RPN: a b & c &

		// Concatenation with quantifiers
		{"ab*", []string{"a", "b", "*", "&"}}, // "a" followed by "b*" -> RPN: a b * &
		{"ab+", []string{"a", "b", "+", "&"}}, // "a" followed by "b+" -> RPN: a b + &
		{"ab?", []string{"a", "b", "?", "&"}}, // "a" followed by "b?" -> RPN: a b ? &

		// Alternation and concatenation
		{"ab|c", []string{"a", "b", "&", "c", "|"}},   // "a" followed by "b" or "c" -> RPN: a b & c |
		{"a(b|c)", []string{"a", "b", "c", "|", "&"}}, // "a" followed by either "b" or "c" -> RPN: a b c | &
		{"a(a+b)*b", []string{"a", "a", "b", "+", "*", "&", "b", "&"}}, //should this be <- or a a b + & * & b &
		// Grouping and concatenation
		{"(ab)c", []string{"a", "b", "&", "c", "&"}},  // "a" followed by "b", then "c" -> RPN: a b & c &
		{"(a|b)c", []string{"a", "b", "|", "c", "&"}}, // Either "a" or "b" followed by "c" -> RPN: a b | c &

		// Concatenation with special characters
		{"a.", []string{"a", ".", "&"}}, // "a" followed by any character -> RPN: a . &
		// {"^ab$", []string{"^", "a", "&", "b", "&", "$", "&"}}, // Start anchor, "a" followed by "b", end anchor -> RPN: ^ a & b & $ &
	}

	for _, test := range tests {
		output := ParseInput(test.input)
		if len(output) != len(test.expected) {
			t.Errorf("ParseInput(%q) = %v; want %v", test.input, output, test.expected)
			continue
		}
		for i := range output {
			if output[i] != test.expected[i] {
				t.Errorf("ParseInput(%q) = %v; want %v", test.input, output, test.expected)
				break
			}
		}
	}
}
