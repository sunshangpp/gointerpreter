package main

import (
//"fmt"
)

type AST interface {
	visit() int
}

type Num struct {
	Token Token
}

type BinOp struct {
	Left  AST
	Right AST
	Token Token
}

type UnaryOp struct {
	Right AST
	Token Token
}

func NewNum(token Token) Num {
	return Num{token}
}

func NewBinOp(left AST, right AST, token Token) BinOp {
	return BinOp{left, right, token}
}

func NewUnaryOp(right AST, token Token) UnaryOp {
	return UnaryOp{right, token}
}

func (num Num) visit() int {
	if num.Token.Type != INTEGER {
		panic("wrong token type")
	}
	return num.Token.Value.(int)
}

func (binOp BinOp) visit() int {
	left := binOp.Left
	right := binOp.Right

	if binOp.Token.Type == PLUS {
		return left.visit() + right.visit()
	} else if binOp.Token.Type == MINUS {
		return left.visit() - right.visit()
	} else if binOp.Token.Type == MUL {
		return left.visit() * right.visit()
	} else if binOp.Token.Type == DIV {
		return left.visit() / right.visit()
	}
	panic("unsupported binOp")
}

func (unaryOp UnaryOp) visit() int {
	if unaryOp.Token.Type == PLUS {
		return unaryOp.Right.visit()
	} else if unaryOp.Token.Type == MINUS {
		return -unaryOp.Right.visit()
	}
	panic("unsupported unaryOp")
}
