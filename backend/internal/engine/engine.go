package engine

// Match checks if the pattern matches anywhere in the text
func Match(pattern string, text string) (bool, []string) {
	return literalMatch([]byte(pattern), []byte(text))
}

// literalMatch checks if the pattern appears anywhere in the text
func literalMatch(pattern []byte, text []byte) (bool, []string) {
	var matches []string
	if len(pattern) == 0 {
		return true, matches
	}

	//literal match with . character
	j := 0
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

	return len(matches) > 0, matches
}
