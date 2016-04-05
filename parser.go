package main

import (
//"fmt"
)

type Parser struct {
	Tokens []Token
	Pos    int
}

func NewParser(tokens []Token) Parser {
	return Parser{Tokens: tokens, Pos: 0}
}

func (parser *Parser) currentToken() Token {
	return parser.Tokens[parser.Pos]
}

func (parser *Parser) eat(tokenTypes ...string) {
	for _, tokenType := range tokenTypes {
		if parser.Tokens[parser.Pos].Type == tokenType {
			parser.Pos += 1
			return
		}
	}
	panic("sytax error!")
}

func (parser *Parser) facter() AST {
	/*
		INTEGER | LPAREN expr RPAREN | (PLUS | MINUS) factor
	*/
	var node AST
	token := parser.currentToken()
	if token.Type == LPAREN {
		parser.eat(LPAREN)
		node = parser.Expr()
		parser.eat(RPAREN)
	} else if token.Type == INTEGER {
		node = NewNum(token)
		parser.eat(INTEGER)
	} else if token.Type == PLUS {
		parser.eat(PLUS)
		node = UnaryOp{parser.facter(), token}
	} else if token.Type == MINUS {
		parser.eat(MINUS)
		node = UnaryOp{parser.facter(), token}
	} else {
		panic("invalid syntax")
	}
	return node
}

func (parser *Parser) term() AST {
	/*
		factor ((MUL | DIV) factor)*
	*/
	node := parser.facter()
	for {
		op := parser.currentToken()
		if op.Type == MUL || op.Type == DIV {
			parser.eat(MUL, DIV)
			node = NewBinOp(node, parser.facter(), op)
		} else {
			return node
		}
	}
}

func (parser *Parser) Expr() AST {
	/*
	   Arithmetic expression parser / interpreter.
	   expr   : term ((PLUS | MINUS) term)*
	   term   : factor ((MUL | DIV) factor)*
	   factor : INTEGER | LPAREN expr RPAREN | (PLUS | MINUS) factor
	*/
	node := parser.term()
	for {
		op := parser.currentToken()
		if op.Type == EOF || op.Type == RPAREN {
			break
		}

		if op.Type == PLUS || op.Type == MINUS {
			parser.eat(PLUS, MINUS)
			node = NewBinOp(node, parser.term(), op)
		} else {
			panic("invalid syntax!")
		}
	}
	return node
}

/*
func (parser *Parser) Expr() {
	result := parser.term()

	for {
		token := parser.currentToken()
		if token.Type == PLUS {
			parser.eat(PLUS)
			result += parser.term()
		} else if token.Type == MINUS {
			parser.eat(MINUS)
			result -= parser.term()
		} else if token.Type == EOF {
			break
		} else {
			parser.eat(PLUS, MINUS)
		}
	}

	fmt.Println(result)
}

func (parser *Parser) StupidExpr() {
	var result int

	previousInt := parser.currentToken()
	parser.eat(INTEGER)
	result = previousInt.Value.(int)

	previousOp := parser.currentToken()
	parser.eat(PLUS, MINUS)

	expect_int := true
	previous_op_plus := false

	if previousOp.Type == PLUS {
		previous_op_plus = true
	}

	for {
		token := parser.currentToken()

		if expect_int {
			parser.eat(INTEGER)
			if previous_op_plus {
				result += token.Value.(int)
			} else {
				result -= token.Value.(int)
			}

			expect_int = false

		} else {
			if token.Type == EOF {
				break
			}
			parser.eat(PLUS, MINUS)
			if token.Type == PLUS {
				previous_op_plus = true
			} else {
				previous_op_plus = false
			}
			expect_int = true
		}

	}

	fmt.Println(result)
}

func (parser *Parser) BasicExpr() {
	left := parser.currentToken()
	parser.eat(INTEGER)

	operand := parser.currentToken()
	if operand.Type == PLUS {
		parser.eat(PLUS)
	} else if operand.Type == MINUS {
		parser.eat(MINUS)
	} else if operand.Type == MULTIPLEX {
		parser.eat(MULTIPLEX)
	} else if operand.Type == DIVIDE {
		parser.eat(DIVIDE)
	}

	right := parser.currentToken()
	parser.eat(INTEGER)

	parser.eat(EOF)

	var result int
	if operand.Type == PLUS {
		result = left.Value.(int) + right.Value.(int)
	} else if operand.Type == MINUS {
		result = left.Value.(int) - right.Value.(int)
	} else if operand.Type == MULTIPLEX {
		result = left.Value.(int) * right.Value.(int)
	} else if operand.Type == DIVIDE {
		result = left.Value.(int) / right.Value.(int)
	}

	fmt.Println(result)
}
*/
