package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func intoPost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	postfix, s := []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)

		case r == ')':
			for s[len(s)-1] != '(' {
				// postfix= append(postfix, s[len(s)-1])
				// s = s[:len(s)-1]    //everything except last char
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = s[:len(s)-1]

		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)

		default:
			postfix = append(postfix, r)
		}
	}

	for len(s) > 0 {
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(postfix)
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

// reader func
func reader() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

func main() {
	// test
	// fmt.Println(pomatch("ab.c*|", "cccc"))

	input := 0
	// read input
	// reader := bufio.NewReader(os.Stdin)

	fmt.Println("Infix / Postfix expression to NFA conversion\n===================")
	// ask for Infix or Postfix
	fmt.Print("Enter:\n1) for Infix \n2) for Postfix\n")

	fmt.Scanln(&input)

	switch input {
	case 1:

		fmt.Println("Enter Infix expression: ")
		// expression, _ := reader.ReadString('\n')
		expression, err := reader()

		// errors
		if err != nil {
			return
		}

		fmt.Println("Enter string to test: ")
		// str, _ := reader.ReadString('\n')
		str, err := reader()

		fmt.Print(expression, str)

		fmt.Println("String", str, " matches nfa: ", pomatch(expression, str))

	case 2:

		// fmt.Println("Enter Postfix expression: ")
		// expression, _ := reader.ReadString('\n')

		// fmt.Println("Enter string to test: ")
		// str, _ := reader.ReadString('\n')

		// fmt.Print(expression, str)

		// fmt.Println("String", str, " matches nfa: ", pomatch(expression, str))

	default:
		fmt.Println("Invalid response!\nPlease enter one of the above: ")
	}

}
