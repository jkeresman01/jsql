package parser

import (
	"fmt"
	"github.com/jkeresman01/jsql/internal/lexer"
	"strings"
)

type Statement struct {
	Type   string
	Table  string
	Values []string
	Name   string
}

type Parser struct {
	l      *lexer.Lexer
	curTok lexer.Token
}

func NewParser(input string) *Parser {
	lex := lexer.NewLexer(input)
	return &Parser{l: lex, curTok: lex.NextToken()}
}

func (p *Parser) next() {
	p.curTok = p.l.NextToken()
}

func (p *Parser) Parse() (*Statement, error) {
	switch p.curTok.Type {
	case lexer.INSERT:
		return p.parseInsert()
	case lexer.SELECT:
		return p.parseSelect()
	case lexer.CREATE:
		return p.parseCreate()
	case lexer.DROP:
		return p.parseDrop()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.curTok)
	}
}

func (p *Parser) parseInsert() (*Statement, error) {
	stmt := &Statement{Type: "INSERT"}

	p.next() // skip INSERT
	if p.curTok.Type != lexer.INTO {
		return nil, fmt.Errorf("expected INTO, got %v", p.curTok.Type)
	}
	p.next()

	if p.curTok.Type != lexer.IDENT {
		return nil, fmt.Errorf("expected table name, got %v", p.curTok.Type)
	}
	stmt.Table = p.curTok.Value
	p.next()

	if p.curTok.Type != lexer.VALUES {
		return nil, fmt.Errorf("expected VALUES, got %v", p.curTok.Type)
	}
	p.next()

	if p.curTok.Type != lexer.LPAREN {
		return nil, fmt.Errorf("expected (, got %v", p.curTok.Type)
	}
	p.next()

	// gather values until ')'
	for p.curTok.Type != lexer.RPAREN && p.curTok.Type != lexer.EOF {
		if p.curTok.Type == lexer.STRING || p.curTok.Type == lexer.NUMBER || p.curTok.Type == lexer.IDENT {
			stmt.Values = append(stmt.Values, p.curTok.Value)
		}
		p.next()
	}

	if p.curTok.Type != lexer.RPAREN {
		return nil, fmt.Errorf("expected closing )")
	}
	return stmt, nil
}

func (p *Parser) parseSelect() (*Statement, error) {
	stmt := &Statement{Type: "SELECT"}
	p.next() // skip SELECT

	if p.curTok.Type != lexer.STAR {
		return nil, fmt.Errorf("only SELECT * supported for now")
	}
	p.next()

	if p.curTok.Type != lexer.FROM {
		return nil, fmt.Errorf("expected FROM, got %v", p.curTok.Type)
	}
	p.next()

	if p.curTok.Type != lexer.IDENT {
		return nil, fmt.Errorf("expected table name after FROM")
	}
	stmt.Table = p.curTok.Value

	return stmt, nil
}

func (p *Parser) parseCreate() (*Statement, error) {
	p.next() // consume CREATE
	if p.curTok.Type != lexer.DATABASE {
		return nil, fmt.Errorf("expected DATABASE keyword")
	}
	p.next()
	if p.curTok.Type != lexer.IDENT {
		return nil, fmt.Errorf("expected database name")
	}
	name := strings.ToLower(p.curTok.Value)
	return &Statement{Type: "CREATE_DATABASE", Name: name}, nil
}

func (p *Parser) parseDrop() (*Statement, error) {
	p.next() // consume DROP
	if p.curTok.Type != lexer.DATABASE {
		return nil, fmt.Errorf("expected DATABASE keyword")
	}
	p.next()
	if p.curTok.Type != lexer.IDENT {
		return nil, fmt.Errorf("expected database name")
	}
	name := strings.ToLower(p.curTok.Value)
	return &Statement{Type: "DROP_DATABASE", Name: name}, nil
}
