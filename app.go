package main

import (
	"fmt"
)

func intopost(infix string) string {
	postfix := ""
	return postfix
}

func main() {

	//ab.c*
	fmt.Println("infix:  ", "a.b.c*")
	fmt.Println("postfix:  ", intopost("a.b.c*"))

}

// type nfa struct {
//     // ...
// }

// func regexcompile(r string) nfa {
//     // ...
//     return n
// }

// func (n nfa) regexmatch(n nfa, r sting) bool {
//     // ...
//     return ismatch
// }
// func main() {
//     n := regexcompile("01*0")
//     t := n.regexmatch("01110")
//     f := n.regexmatch("1000001")
// }
