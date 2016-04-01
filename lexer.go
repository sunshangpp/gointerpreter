package main

import (
	"strconv"
)

const (
	EOF     = "EOF"
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	MUL     = "MUL"
	DIV     = "DIV"
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
)

type Token struct {
	Type  string
	Value interface{}
}

type Lexer struct {
	Text        string
	Pos         int
	CurrentChar string
}

func NewLexer(text string) Lexer {
	var currentChar string
	if len(text) == 0 {
		currentChar = EOF
	} else {
		currentChar = text[0:1]
	}
	return Lexer{Text: text, Pos: 0, CurrentChar: currentChar}
}

func (lexer *Lexer) Tokenize() []Token {
	tokens := []Token{}

	for {
		token := lexer.nextToken()
		tokens = append(tokens, token)
		if token.Type == EOF {
			break
		}
	}
	return tokens
}

// currently only supports multi-digit integer, plus and minus; space allowed
func (lexer *Lexer) nextToken() Token {
	lexer.skipSpace()

	if lexer.CurrentChar == EOF {
		return Token{Type: EOF, Value: EOF}
	} else if isDigit, _ := isDigit(lexer.CurrentChar); isDigit == true {
		return Token{Type: INTEGER, Value: lexer.getInteger()}
	} else if lexer.CurrentChar == "+" {
		lexer.advance()
		return Token{Type: PLUS, Value: "+"}
	} else if lexer.CurrentChar == "-" {
		lexer.advance()
		return Token{Type: MINUS, Value: "-"}
	} else if lexer.CurrentChar == "-" {
		lexer.advance()
		return Token{Type: MINUS, Value: "-"}
	} else if lexer.CurrentChar == "*" {
		lexer.advance()
		return Token{Type: MUL, Value: "*"}
	} else if lexer.CurrentChar == "/" {
		lexer.advance()
		return Token{Type: DIV, Value: "/"}
	} else if lexer.CurrentChar == "(" {
		lexer.advance()
		return Token{Type: LPAREN, Value: "("}
	} else if lexer.CurrentChar == ")" {
		lexer.advance()
		return Token{Type: RPAREN, Value: ")"}
	}

	panic("unsupported character!")
}

func (lexer *Lexer) advance() {
	lexer.Pos += 1
	if lexer.Pos >= len(lexer.Text) {
		lexer.CurrentChar = EOF
	} else {
		lexer.CurrentChar = lexer.Text[lexer.Pos : lexer.Pos+1]
	}
}

func isDigit(char string) (bool, int) {
	value, err := strconv.Atoi(char)
	if err != nil {
		return false, -1
	}
	return true, value
}

func (lexer *Lexer) skipSpace() {
	for lexer.CurrentChar == " " {
		lexer.advance()
	}
}

func (lexer *Lexer) getInteger() int {
	value := 0
	for lexer.CurrentChar != EOF {
		isDigit, digit := isDigit(lexer.CurrentChar)

		if isDigit == true {
			value = value*10 + digit
			lexer.advance()
		} else {
			break
		}
	}
	return value
}
