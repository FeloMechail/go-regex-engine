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

func (s *State) AddTransition(char rune, state *State) {
	s.Transitions[char] = append(s.Transitions[char], state)
}

func (s *State) AddEpsilonTransition(state *State) {
	s.epsilonTransitions = append(s.epsilonTransitions, state)
}

func NewNFA(start, accept *State) *NFA {
	return &NFA{
		Start:  start,
		Accept: accept,
	}
}
