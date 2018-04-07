package main

import (

	"bufio"
	"fmt"
	"os"

	nfa "./nfa"
)

func main(){

	option := 0
	quit := true


	for quit {

		//User prompt
		fmt.Println("Please choose from the following\n")
		fmt.Println("Infix to Postfix\n")
		fmt.Println("Postfix to NFA\n")
		fmt.Println("Check if String equals a regular Expression")
		fmt.Scanln(&option)

		if option == 1 {

			// Reading in expression
			fmt.Println("Please enter an expression")
			read := bufio.NewReader(os.Stdio)
			exp, _ := read.ReadString('\n')
			exp = nfa.StringTirm(exp)

			fmt.Println("Infix: ", exp)
			fmt.Println("Postfix: ", nfa.Intopost(exp))
			

			

		}




		

		
	}
	
}