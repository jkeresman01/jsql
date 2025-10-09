package db

import "fmt"

type Statement struct {
	Type   string
	Table  string
	Values []string
}

type Parser struct {
	l      *Lexer
	curTok Token
}

func NewParser(input string) *Parser {
	lex := NewLexer(input)
	return &Parser{l: lex, curTok: lex.NextToken()}
}

func (p *Parser) next() {
	p.curTok = p.l.NextToken()
}

func (p *Parser) Parse() (*Statement, error) {
	switch p.curTok.Type {
	case TOK_INSERT:
		return p.parseInsert()
	case TOK_SELECT:
		return p.parseSelect()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.curTok)
	}
}

func (p *Parser) parseInsert() (*Statement, error) {
	stmt := &Statement{Type: "INSERT"}

	p.next() // skip INSERT
	if p.curTok.Type != TOK_INTO {
		return nil, fmt.Errorf("expected INTO, got %v", p.curTok.Type)
	}
	p.next()

	if p.curTok.Type != TOK_IDENT {
		return nil, fmt.Errorf("expected table name, got %v", p.curTok.Type)
	}
	stmt.Table = p.curTok.Value
	p.next()

	if p.curTok.Type != TOK_VALUES {
		return nil, fmt.Errorf("expected VALUES, got %v", p.curTok.Type)
	}
	p.next()

	if p.curTok.Type != TOK_LPAREN {
		return nil, fmt.Errorf("expected (, got %v", p.curTok.Type)
	}
	p.next()

	// gather values until ')'
	for p.curTok.Type != TOK_RPAREN && p.curTok.Type != TOK_EOF {
		if p.curTok.Type == TOK_STRING || p.curTok.Type == TOK_NUMBER || p.curTok.Type == TOK_IDENT {
			stmt.Values = append(stmt.Values, p.curTok.Value)
		}
		p.next()
	}

	if p.curTok.Type != TOK_RPAREN {
		return nil, fmt.Errorf("expected closing )")
	}
	return stmt, nil
}

func (p *Parser) parseSelect() (*Statement, error) {
	stmt := &Statement{Type: "SELECT"}
	p.next() // skip SELECT

	if p.curTok.Type != TOK_STAR {
		return nil, fmt.Errorf("only SELECT * supported for now")
	}
	p.next()

	if p.curTok.Type != TOK_FROM {
		return nil, fmt.Errorf("expected FROM, got %v", p.curTok.Type)
	}
	p.next()

	if p.curTok.Type != TOK_IDENT {
		return nil, fmt.Errorf("expected table name after FROM")
	}
	stmt.Table = p.curTok.Value

	return stmt, nil
}
