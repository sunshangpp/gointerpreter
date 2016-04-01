package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("calc > ")

	for scanner.Scan() {
		line := scanner.Text()

		// lexecal analysis: tokenize input
		lexer := NewLexer(line)
		tokens := lexer.Tokenize()
		fmt.Println(tokens)

		// parse tokens
		parser := NewParser(tokens)
		result := parser.Expr()
		fmt.Println(result)

		fmt.Print("calc > ")
	}
}
