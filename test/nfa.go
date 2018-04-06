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

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			// push to stack
			nfastack = append(nfastack, &nfa{initial: &accept, accept: &accept})

		}
	}

	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
