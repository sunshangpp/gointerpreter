package main

import (
	"fmt"
)

type Interpreter struct {
	Text string
}

func NewInterpreter(text string) Interpreter {
	return Interpreter{text}
}

func (interpreter *Interpreter) interpret() int {
	lexer := NewLexer(interpreter.Text)
	tokens := lexer.Tokenize()
	fmt.Println(tokens)

	// parse tokens
	parser := NewParser(tokens)
	ast := parser.Expr()
	return ast.visit()
}
