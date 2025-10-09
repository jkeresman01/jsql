package db

import (
	"strings"
	"unicode"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input []rune
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: []rune(input)}
}

func (l *Lexer) NextToken() Token {
	// skip whitespace
	for l.pos < len(l.input) && unicode.IsSpace(l.input[l.pos]) {
		l.pos++
	}

	if l.pos >= len(l.input) {
		return Token{Type: TOK_EOF}
	}

	ch := l.input[l.pos]
	l.pos++

	switch ch {
	case '(':
		return Token{Type: TOK_LPAREN, Value: string(ch)}
	case ')':
		return Token{Type: TOK_RPAREN, Value: string(ch)}
	case ',':
		return Token{Type: TOK_COMMA, Value: string(ch)}
	case ';':
		return Token{Type: TOK_SEMI, Value: string(ch)}
	case '*':
		return Token{Type: TOK_STAR, Value: string(ch)}
	case '\'', '"':
		start := ch
		var val []rune
		for l.pos < len(l.input) && l.input[l.pos] != start {
			val = append(val, l.input[l.pos])
			l.pos++
		}
		l.pos++ // skip closing quote
		return Token{Type: TOK_STRING, Value: string(val)}
	}

	// identifiers / keywords / numbers
	start := l.pos - 1
	for l.pos < len(l.input) &&
		(unicode.IsLetter(l.input[l.pos]) ||
			unicode.IsDigit(l.input[l.pos]) ||
			l.input[l.pos] == '_') {
		l.pos++
	}
	word := string(l.input[start:l.pos])
	up := strings.ToUpper(word)

	switch up {
	case "INSERT":
		return Token{Type: TOK_INSERT, Value: up}
	case "INTO":
		return Token{Type: TOK_INTO, Value: up}
	case "VALUES":
		return Token{Type: TOK_VALUES, Value: up}
	case "SELECT":
		return Token{Type: TOK_SELECT, Value: up}
	case "FROM":
		return Token{Type: TOK_FROM, Value: up}
	default:
		if unicode.IsDigit(rune(word[0])) {
			return Token{Type: TOK_NUMBER, Value: word}
		}
		return Token{Type: TOK_IDENT, Value: word}
	}
}

