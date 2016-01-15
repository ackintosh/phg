package main

import (
	"regexp"
	"text/scanner"
)

var keywords = map[string]int{
	"echo": ECHO,
}

type Lexer struct {
	scanner.Scanner
	result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	literal := l.TokenText()
	if token == scanner.Int {
		token = NUMBER
	}
	if token == scanner.Ident {
		token = STRING
	}
	r := regexp.MustCompile(`\".*\"`)
	if r.MatchString(literal) {
		token = STRING
	}
	if keyword, ok := keywords[literal]; ok {
		token = keyword
	}

	lval.token = Token{token: token, literal: literal}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}
