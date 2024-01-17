package token

// helps define multiple types of tokens
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Help define tokens
const (
	NA  = "NA"  // not available
	EOF = "EOF" // end of file

	//Identifiers + literals
	IDENT  = "IDENT"  // add, x, y, ..
	INT    = "INT"    // 1234567890
	FLOAT  = "FLOAT"  // 1.0,1.2,2.3 etc
	STRING = "STRING" // ""
	CHAR   = "CHAR"   // ''
	NULL   = "NULL"   // null

	//Operators
	ASSIGN = "="
	PLUS   = "+"
	SUB    = "-"
	DIV    = "/"
	MUL    = "*"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACK    = "["
	RBRACK    = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
