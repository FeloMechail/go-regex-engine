package engine

import (
	"testing"
)

func TestNewState(t *testing.T) {
	state := NewState()

	if state == nil {
		t.Fatalf("Expected new state to be non-nil")
	}

	if state.Transitions == nil {
		t.Errorf("Expected Transitions to be initialized, got nil")
	}

	if len(state.Transitions) != 0 {
		t.Errorf("Expected Transitions to be empty, got %d elements", len(state.Transitions))
	}

	if state.epsilonTransitions == nil {
		t.Errorf("Expected epsilonTransitions to be initialized, got nil")
	}

	if len(state.epsilonTransitions) != 0 {
		t.Errorf("Expected epsilonTransitions to be empty, got %d elements", len(state.epsilonTransitions))
	}

	if state.Accepting != false {
		t.Errorf("Expected Accepting to be false, got %v", state.Accepting)
	}
}

func TestAddTransition(t *testing.T) {
	state1 := NewState()
	state2 := NewState()

	state1.AddTransition('a', state2)

	if len(state1.Transitions) != 1 {
		t.Fatalf("Expected 1 transition, got %d", len(state1.Transitions))
	}

	if states, ok := state1.Transitions['a']; !ok {
		t.Fatalf("Expected transition on 'a' to be present")
	} else {
		if len(states) != 1 {
			t.Fatalf("Expected 1 state in transition list, got %d", len(states))
		}

		if states[0] != state2 {
			t.Fatalf("Expected state2 in transition list, got different state")
		}
	}
}

func TestNewNFA(t *testing.T) {
	startState := NewState()
	acceptState := NewState()

	nfa := NewNFA(startState, acceptState)

	if nfa == nil {
		t.Fatalf("Expected new NFA to be non-nil")
	}

	if nfa.Start != startState {
		t.Errorf("Expected Start state to be %v, got %v", startState, nfa.Start)
	}

	if nfa.Accept != acceptState {
		t.Errorf("Expected Accept state to be %v, got %v", acceptState, nfa.Accept)
	}
}
