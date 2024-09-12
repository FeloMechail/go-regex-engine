package engine

import "bytes"

// Match checks if the pattern matches anywhere in the text
func Match(pattern string, text string) bool {
	return literalMatch([]byte(pattern), []byte(text))
}

// literalMatch checks if the pattern appears anywhere in the text
func literalMatch(pattern []byte, text []byte) bool {
	if len(pattern) == 0 {
		return true
	}

	// Use bytes.Contains to check if the pattern exists in the text
	return bytes.Contains(text, pattern)
}
