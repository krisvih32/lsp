package main

import (
	"fmt"
	"strings"
)

func lexer(code string) (toknumlist []int) {
	// print("in lex")
	tokens := strings.Split(code, " ")
	// fmt.Printf("tokens %v", tokens)
	for i := 0; i < len(tokens); i++ {
		strings.Replace(tokens[i], "  ", " ", -1)
		// print("in for")
		switch tokens[i] {
		case "call":
			toknumlist = append(toknumlist, tokencall)
		case "ret":
			toknumlist = append(toknumlist, tokenret)
		case "":
		default:
			printerror(fmt.Sprintf("hello. in token \"%s\" token number %d, you made a mistake", tokens[i], i))
			// printerror(fmt.Sprintf("%v", tokens))
		}
	}
	return
}
