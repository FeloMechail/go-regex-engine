package regex

import "regex-engine/internal/engine"

// Match checks if the pattern matches anywhere in the text
func Match(pattern string, text string) (bool, []string) {
	return engine.Match(pattern, text)
}
