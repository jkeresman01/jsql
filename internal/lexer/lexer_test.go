package lexer

import (
	"testing"
)

func TestSimpleInsertTokens(t *testing.T) {
	input := "INSERT INTO users VALUES (1, 'Alice');"
	l := NewLexer(input)

	expected := []Token{
		{Type: INSERT, Value: "INSERT"},
		{Type: INTO, Value: "INTO"},
		{Type: IDENT, Value: "users"},
		{Type: VALUES, Value: "VALUES"},
		{Type: LPAREN, Value: "("},
		{Type: NUMBER, Value: "1"},
		{Type: COMMA, Value: ","},
		{Type: STRING, Value: "Alice"},
		{Type: RPAREN, Value: ")"},
		{Type: SEMI, Value: ";"},
		{Type: EOF, Value: ""},
	}

	for i, expectedTok := range expected {
		actual := l.NextToken()
		if actual.Type != expectedTok.Type {
			t.Fatalf("token[%d] - expected type %v, got %v", i, expectedTok.Type, actual.Type)
		}
		if actual.Value != expectedTok.Value {
			t.Fatalf("token[%d] - expected value %q, got %q", i, expectedTok.Value, actual.Value)
		}
	}
}

func TestSelectTokens(t *testing.T) {
	input := "SELECT * FROM users;"
	l := NewLexer(input)

	expected := []TokenType{
		SELECT, STAR, FROM, IDENT, SEMI, EOF,
	}

	for i, typ := range expected {
		actual := l.NextToken()
		if actual.Type != typ {
			t.Fatalf("token[%d] - expected type %v, got %v", i, typ, actual.Type)
		}
	}
}
