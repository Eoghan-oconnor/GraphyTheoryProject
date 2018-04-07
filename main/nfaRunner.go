package main

import (

	"bufio"
	"fmt"
	"os"

	nfa ".."
)

func main(){

	option := 0
	quit := true


	for quit {

		//User prompt
		fmt.Println("Please choose from the following\n")
		fmt.Println("======================================")
		fmt.Println("1) Infix to Postfix\n")
		fmt.Println("2) Postfix to NFA\n")
		fmt.Scanln(&option)

		if option == 1 {

			// Reading in expression
			fmt.Println("Please enter an expression")
			read := bufio.NewReader(os.Stdin)
			exp, _ := read.ReadString('\n')
			exp = nfa.StringTrim(exp)

			//outputting input and result
			fmt.Println("Infix: ", exp)
			fmt.Println("Postfix: ", nfa.intopost(exp))
		
		} else if option == 2 {

			//Reading in expression
			fmt.Println("Enter an expression to be converted")
			read := bufio.NewReader(os.Stdin)
			exp, _ := read.ReadString('\n')
			exp = nfa.StringTrim(exp)

			// Outputting input and result
			fmt.Println("Postfix: ", exp)
			fmt.Println("NFA: ", nfa.poregtonfa(exp))
			
			// Enter String 
			fmt.Print("Enter a String to test againist NFA: ")
			s, _ := read.ReadString('\n')
			s = nfa.StringTrim(s)
			s = nfa.intopost(s)

			if nfa.pomatch(exp, s) == true {

				fmt.Println("String Matches expression\n")
			} else {
				fmt.Println("String does not match expression\n")
			}


		}





		

		
	}
	
}