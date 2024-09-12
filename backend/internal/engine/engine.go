package engine

import "log"

// Match checks if the pattern matches anywhere in the text
func Match(pattern string, text string) bool {
	return literalMatch([]byte(pattern), []byte(text))
}

// literalMatch checks if the pattern appears anywhere in the text
func literalMatch(pattern []byte, text []byte) bool {
	if len(pattern) == 0 {
		return true
	}

	//literal match with . character
	j := 0
	var matches []string
	for i := 0; i < len(text); i++ {
		if pattern[j] == text[i] || (pattern[j] == '.' && text[i] != '\n') {
			j++
		} else {
			j = 0
		}

		if j == len(pattern) {
			matches = append(matches, string(text[i-len(pattern)+1:i+1]))
			j = 0
		}
	}

	log.Println(matches)

	return len(matches) > 0
}
