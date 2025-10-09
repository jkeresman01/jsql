package lexer

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	// Keywords
	INSERT TokenType = "INSERT"
	INTO   TokenType = "INTO"
	VALUES TokenType = "VALUES"
	SELECT TokenType = "SELECT"
	FROM   TokenType = "FROM"

	// Symbols
	STAR   TokenType = "*"
	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	COMMA  TokenType = ","
	SEMI   TokenType = ";"

	// Literals
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"
	IDENT  TokenType = "IDENT"

	EOF TokenType = "EOF"
)

