package db

type TokenType string

const (
	// keywords
	TOK_INSERT TokenType = "INSERT"
	TOK_INTO   TokenType = "INTO"
	TOK_VALUES TokenType = "VALUES"
	TOK_SELECT TokenType = "SELECT"
	TOK_FROM   TokenType = "FROM"

	// symbols
	TOK_STAR    TokenType = "*"
	TOK_LPAREN  TokenType = "("
	TOK_RPAREN  TokenType = ")"
	TOK_COMMA   TokenType = ","
	TOK_SEMI    TokenType = ";"
	TOK_STRING  TokenType = "STRING"
	TOK_NUMBER  TokenType = "NUMBER"
	TOK_IDENT   TokenType = "IDENT"
	TOK_EOF     TokenType = "EOF"
)

