package parser

import (
	"github.com/kenju/go-minruby/ast"
	"github.com/kenju/go-minruby/lexer"
)

type Parser struct {
	l *lexer.Lexer
	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	// TODO: iterates over every token until EOF

	return program
}
