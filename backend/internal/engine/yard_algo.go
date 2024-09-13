package engine

import "fmt"

var precedence = map[rune]struct {
	prec      int
	righAssoc bool
}{
	'^': {4, true},
	'*': {3, false},
	'/': {3, false},
	'+': {2, false},
	'-': {2, false},
}

//testInput := "3+4*2(1-5)*2*3"

func ParseInput(input string) (rpn string) {
	stack := Stack{}
	for _, char := range input {
		switch char {
		case '(':
			stack.Push(char)
		case ')':
			if op, err := stack.Top(); err != nil {
				fmt.Println(err)
			} else {
				for op != '(' {
					rpn += " " + string(op)
					stack.Pop()
					op, _ = stack.Top()
				}
				//pop final (
				stack.Pop()
			}
		default:
			if prec, isOp := precedence[char]; isOp {
				for !stack.IsEmpty() {
					op, _ := stack.Top()
					if prec2, isOp := precedence[op]; !isOp || prec.prec > prec2.prec || prec.prec == prec2.prec && prec.righAssoc {
						break
					}
					//top item is an operator
					stack.Pop()
					rpn += " " + string(op)
				}
				stack.Push(char)
			} else {
				if rpn > "" {
					rpn += " "
				}
				rpn += string(char)
			}
		}
	}

	for !stack.IsEmpty() {
		op, _ := stack.Top()
		rpn += " " + string(op)
		stack.Pop()
	}
	return
}
