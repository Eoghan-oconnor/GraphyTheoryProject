package main

import (
	"fmt"
)

func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	pofix := []rune{}
	s := []rune{}

}

func main() {

	//Answear ab.c*
	fmt.Println("infix:	  ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer abd|.*
	fmt.Println("Infix:  ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//Answer abd|.c*
	fmt.Println("Infix:  ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	//Answer abb.*.c
	fmt.Println("Infix:  ", "a.(b.b)*.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)*.c"))
}
