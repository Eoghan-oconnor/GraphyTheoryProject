package nfa

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func Poregtonfa(pofix string) *nfa {
	nfaStack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.accept}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '?':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag.initial, edge2: frag.accept}

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: frag.accept})
		case '+':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = &initial

			nfaStack = append(nfaStack, &nfa{initial: frag.initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		} //end switch
	} //end for

	return nfaStack[0]
}

//Add state function
//This function returns an array of pointers back to the state Struct
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
} //End addState function

func Pomatch(po string, s string) bool {
	ismatch := false

	ponfa := Poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			} //End if
		} //End for
		current, next = next, []*state{}
	} //End for

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		} //End if

	} //End range current for

	//return true/false result
	return ismatch
}

//	main() {
// 	nfa := poregtonfa("ab.c*|", "ccc")
// 	fmt.Println(nfa)
// }

//
