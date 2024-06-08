package lexer

import (
	"skibidilang/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	nextPosition int  // next position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	//delimiters
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
		//operators
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewTwoCharToken(token.EQ, "==")
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '+':
		if l.peekChar() == '+' {
			l.readChar()
			tok = token.NewTwoCharToken(token.INC, "++")
		} else {
			tok = token.NewToken(token.ADD, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			l.readChar()
			tok = token.NewTwoCharToken(token.DEC, "--")
		} else {
			tok = token.NewToken(token.SUB, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewTwoCharToken(token.NEQ, "!=")
		} else {
			tok = token.NewToken(token.NOT, l.ch)
		}
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '^':
		tok = token.NewToken(token.POWER, l.ch)
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewTwoCharToken(token.LEQ, "<=")
		} else {
			tok = token.NewToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewTwoCharToken(token.GEQ, ">=")
		} else {
			tok = token.NewToken(token.GT, l.ch)
		}
	case '&':
		tok = token.NewToken(token.AND, l.ch)
	case '|':
		tok = token.NewToken(token.OR, l.ch)
	case '%':
		tok = token.NewToken(token.MOD, l.ch)
		//brackets
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '[':
		tok = token.NewToken(token.LBRACK, l.ch)
	case ']':
		tok = token.NewToken(token.RBRACK, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}
