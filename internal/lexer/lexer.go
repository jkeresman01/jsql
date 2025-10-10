package lexer

import (
	"strings"
	"unicode"
)

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
		return Token{Type: EOF}
	}

	ch := l.input[l.pos]
	l.pos++

	switch ch {
	case '(':
		return Token{Type: LPAREN, Value: string(ch)}
	case ')':
		return Token{Type: RPAREN, Value: string(ch)}
	case ',':
		return Token{Type: COMMA, Value: string(ch)}
	case ';':
		return Token{Type: SEMI, Value: string(ch)}
	case '*':
		return Token{Type: STAR, Value: string(ch)}
	case '\'', '"':
		start := ch
		var val []rune
		for l.pos < len(l.input) && l.input[l.pos] != start {
			val = append(val, l.input[l.pos])
			l.pos++
		}
		l.pos++ // skip closing quote
		return Token{Type: STRING, Value: string(val)}
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
		return Token{Type: INSERT, Value: up}
	case "INTO":
		return Token{Type: INTO, Value: up}
	case "VALUES":
		return Token{Type: VALUES, Value: up}
	case "SELECT":
		return Token{Type: SELECT, Value: up}
	case "FROM":
		return Token{Type: FROM, Value: up}
	case "CREATE":
		return Token{Type: CREATE, Value: up}
	case "DROP":
		return Token{Type: DROP, Value: up}
	case "DATABASE":
		return Token{Type: DATABASE, Value: up}

	default:
		if unicode.IsDigit(rune(word[0])) {
			return Token{Type: NUMBER, Value: word}
		}
		return Token{Type: IDENT, Value: word}
	}
}
