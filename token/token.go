package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN   = "="
	ADD      = "+"
	SUB      = "-"
	NOT      = "!"
	ASTERISK = "*"
	SLASH    = "/"
	AND      = "&"
	OR       = "|"
	MOD      = "%"
	LT       = "<"
	GT       = ">"
	POWER    = "^"
	//Two char tokens
	EQ  = "=="
	NEQ = "!="
	INC = "++"
	DEC = "--"
	LEQ = "<="
	GEQ = ">="
	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"ohio":    FUNCTION,
	"skibidi": LET,
	"alpha":   TRUE,
	"beta":    FALSE,
	"if":      IF,
	"else":    ELSE,
	"goon":    RETURN,
}

// LookupIdent return token type based on an identifier string
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func NewTwoCharToken(tokenType TokenType, s string) Token {
	return Token{Type: tokenType, Literal: s}
}
