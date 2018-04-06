package main

func intopost(infix string) string {
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

// func main() {

// 	//ab.c*
// 	fmt.Println("infix:  ", "a.b.c*")
// 	fmt.Println("postfix:  ", intopost("a.b.c*"))
// 	//abd|*
// 	fmt.Println("infix:  ", "(a.(b|d))*")
// 	fmt.Println("postfix:  ", intopost("(a.(b|d))*"))
// 	//abd|.c*
// 	fmt.Println("infix:  ", "a.(b|d).c*")
// 	fmt.Println("postfix:  ", intopost("a.(b|d).c*"))
// 	//abb.+.c.
// 	fmt.Println("infix:  ", "a.(b.b)+.c")
// 	fmt.Println("postfix:  ", intopost("a.(b.b)+.c"))

// }
