package engine

type NFA struct {
	Start  *State
	Accept *State
}

type State struct {
	Transitions        map[rune][]*State
	epsilonTransitions []*State
	Accepting          bool
}

func NewState() *State {
	return &State{
		Transitions:        make(map[rune][]*State),
		epsilonTransitions: make([]*State, 0),
		Accepting:          false,
	}
}

func (s *State) AddTransition(char rune, next *State) {
	s.Transitions[char] = append(s.Transitions[char], next)
}

func (s *State) AddEpsilonTransition(next *State) {
	s.epsilonTransitions = append(s.epsilonTransitions, next)
}

func BuildNFA(start, accept *State) *NFA {
	return &NFA{
		Start:  start,
		Accept: accept,
	}
}

func EpsilonNFA() *NFA {
	q0 := NewState()
	q1 := NewState()
	q1.Accepting = true
	q0.AddEpsilonTransition(q1)

	return BuildNFA(q0, q1)
}

func LiteralNFA(char rune) *NFA {
	q0 := NewState()
	q1 := NewState()
	q1.Accepting = true
	q0.AddTransition(char, q1)

	return BuildNFA(q0, q1)
}

func concatFNA(first, second *NFA) *NFA {
	first.Accept.Accepting = false
	first.Accept.AddEpsilonTransition(second.Start)

	return BuildNFA(first.Start, second.Accept)
}

func unionNFA(first, second *NFA) *NFA {
	i := NewState()
	f := NewState()
	f.Accepting = true

	i.AddEpsilonTransition(first.Start)
	i.AddEpsilonTransition(second.Start)

	first.Accept.Accepting = false
	second.Accept.Accepting = false

	first.Accept.AddEpsilonTransition(f)
	second.Accept.AddEpsilonTransition(f)

	return BuildNFA(i, f)
}

func closureNFA(first *NFA) *NFA {
	i := NewState()
	f := NewState()
	f.Accepting = true

	i.AddEpsilonTransition(f)
	i.AddEpsilonTransition(first.Start)
	first.Accept.AddEpsilonTransition(f)
	first.Accept.AddEpsilonTransition(first.Start)
	first.Accept.Accepting = false

	return BuildNFA(i, f)
}

func ConstructNFA(rpn []rune) *NFA {
	if len(rpn) == 0 {
		return EpsilonNFA()
	}
	stack := Stack[*NFA]{}

	for _, char := range rpn {
		switch char {
		case '*':
			char, _ := stack.Top()
			stack.Push(closureNFA(char))
			stack.Pop()
		case '|':
			right, _ := stack.Top()
			stack.Pop()
			left, _ := stack.Top()
			stack.Pop()
			stack.Push(unionNFA(left, right))
		case '&':
			right, _ := stack.Top()
			stack.Pop()
			left, _ := stack.Top()
			stack.Pop()
			stack.Push(concatFNA(left, right))
		default:
			stack.Push(LiteralNFA(char))
		}
	}

	finalNFA, _ := stack.Top()
	return finalNFA
}
