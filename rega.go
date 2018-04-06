package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func poregtonfa(posfix string) *nfa {
	nfastack := []*nfa{}

	for _, r := range posfix {
		switch r {
		case '.':
			// first is at bottom
			frag2 := nfastack[len(nfastack)-1] // pop last off stack
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// concat accept state of first to initial of second
			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept}) // frag2 accept state is accept state of new fragment
		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			// new initial state
			initial := state{edge1: frag1.initial, edge2: frag2.initial}

			// join accept states to new state
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

			////
			fmt.Println(" ", initial)

		case '*':
			// pop one frag off stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			// push new frag
			// old frag with new accept and initial states
			nfastack = append(nfastack, &nfa{initial: &accept, accept: &accept})

			////
			fmt.Println(" ", initial)

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			// push to stack
			nfastack = append(nfastack, &nfa{initial: &accept, accept: &accept})
			////
			fmt.Println(" ", initial)

		}
	}

	if len(nfastack) != 1 {
		fmt.Println("Oops: ", len(nfastack), nfastack)
	}

	return nfastack[0]
}

func pomatch(po string, s string) bool {
	ismatch := false

	ponfa := poregtonfa(po)

	// lists of states
	current := []*state{ponfa.initial}
	next := []*state{}

	// add initial state of postfix nfa to current array
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// generate next from current
	for _, r := range s {
		for _, c := range current {
			// check if labeled
			if c.symbol == r {
				// add c state to array
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// set current to next
		current, next = next, []*state{}
	}

	// loop and check if accept state
	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
		}
	}
	return ismatch
}

func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 { // not in accept state
		l = addState(l, s.edge1, a)
		if s.edge2 != nil { // if there is second edge
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

func main() {
	fmt.Println(pomatch("ab.c*|", "cccc"))
}
