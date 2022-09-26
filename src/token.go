package main

type Token struct {
	tokenType TokenType
	lexeme string
	literal interface{}
	line int
}