package engine

type Matches struct {
	Match  string
	Rangee []int
}

// Match checks if the pattern matches anywhere in the text
func Match(pattern string, text string) (bool, []Matches) {
	return literalMatch([]byte(pattern), []byte(text))
}

// literalMatch checks if the pattern appears anywhere in the text
func literalMatch(pattern []byte, text []byte) (bool, []Matches) {

	found := make([]Matches, 0)

	if len(pattern) == 0 {
		return true, found
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
			found = append(found, Matches{string(text[i-len(pattern)+1 : i+1]), []int{i - len(pattern) + 1, i + 1}})
			j = 0
		}
	}

	return len(found) > 0, found
}
