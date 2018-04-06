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

		case '|':

		case '*':

		default:

		}
	}

	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
