package regex_test

import (
	"regex-engine/pkg/regex"
	"testing"
)

func TestLiteralMatch(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		input   string
		want    bool
	}{
		{"ExactMatch", "hello", "hello", true},
		{"PartialMatch", "hello", "hello world", true},
		{"NoMatch", "hello", "world", false},
		{"EmptyPattern", "", "hello", true},
		{"EmptyInput", "hello", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := regex.Match(tt.pattern, tt.input)
			if got != tt.want {
				t.Errorf(
					"Match(%q, %q) = %v, want %v",
					tt.pattern,
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestBasicCharacterMatching(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		input   string
		want    bool
	}{
		{"SingleChar", "a", "a", true},
		{"SingleCharNoMatch", "a", "b", false},
		{"AnyChar", ".", "x", true},
		{"AnyCharNewline", ".", "\n", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := regex.Match(tt.pattern, tt.input)
			if got != tt.want {
				t.Errorf(
					"Match(%q, %q) = %v, want %v",
					tt.pattern,
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}

// Add more test functions for each feature as you implement them
// For example:
// func TestCharacterClasses(t *testing.T) { ... }
// func TestQuantifiers(t *testing.T) { ... }
// func TestAlternation(t *testing.T) { ... }
// func TestGrouping(t *testing.T) { ... }
// func TestBacktracking(t *testing.T) { ... }
// func TestAnchors(t *testing.T) { ... }
// func TestEscapeCharacters(t *testing.T) { ... }
