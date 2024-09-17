package engine

import (
	"fmt"
)

// var precedence = map[rune]struct {
// 	prec      int
// 	righAssoc bool
// }{
// 	'^': {4, true},
// 	'*': {3, false},
// 	'/': {3, false},
// 	'+': {2, false},
// 	'-': {2, false},
// }

var regexPrecedence = map[rune]struct {
	prec      int
	righAssoc bool
}{
	'*': {5, false},
	'+': {5, false},
	'?': {5, false},
	'&': {4, false},
	'|': {3, true},
	'^': {1, false},
	'$': {1, false},
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// FIXME: concat after + if (a+b), maybe dont need fixing
func addConcatenation(input string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i])
		if i+1 < len(input) {
			current := input[i]
			next := input[i+1]

			if (isLetter(current) || current == ')' || current == '*' || current == '+' || current == '?') &&
				(isLetter(next) || next == '(' || next == '.') {
				output += "&"
			}
		}
	}
	return
}

func ParseInput(input string) (rpn []rune) {
	stack := Stack[rune]{}
	input = addConcatenation(input)
	for _, char := range input {
		switch char {
		case ' ':
			continue
		case '(':
			stack.Push(char)
		case ')':
			if op, err := stack.Top(); err != nil {
				fmt.Println(err)
			} else {
				for op != '(' {
					// rpn += " " + string(op)
					rpn = append(rpn, op)
					stack.Pop()
					op, _ = stack.Top()
				}
				// pop final (
				stack.Pop()
			}
		default:
			if prec, isOp := regexPrecedence[char]; isOp {
				for !stack.IsEmpty() {
					op, _ := stack.Top()
					if prec2, isOp := regexPrecedence[op]; !isOp ||
						prec.prec > prec2.prec ||
						prec.prec == prec2.prec && prec.righAssoc {
						break
					}
					// top item is an operator
					stack.Pop()
					// rpn += " " + string(op)
					rpn = append(rpn, op)
				}
				stack.Push(char)
			} else {
				// if rpn > "" {
				// 	rpn += " "
				// }
				// rpn += string(char)
				rpn = append(rpn, char)
			}
		}
	}

	for !stack.IsEmpty() {
		op, _ := stack.Top()
		// rpn += " " + string(op)
		rpn = append(rpn, op)
		stack.Pop()
	}
	return
}
