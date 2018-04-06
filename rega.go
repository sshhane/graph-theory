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

func postRegexNfa(posfix string) *nfa {
	nfaStack := []*nfa{}

	for _, r := range posfix {
		switch r {
		case '.':
			// first is at bottom
			frag2 := nfaStack[len(nfaStack)-1] // pop last off stack
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			// concat accept state of first to initial of second
			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept}) // frag2 accept state is accept state of new fragment

		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			// new initial state
			initial := state{edge1: frag1.initial, edge2: frag2.initial}

			// join accept states to new state
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '*':
			// pop one frag off stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			// push new frag
			// old frag with new accept and initial states
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			// push to stack
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		}
	}

	return nfaStack[0]
}

func pomatch(po string, s string) bool {

	ismatch := false

	ponfa := postRegexNfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}

		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
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
