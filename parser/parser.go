package parser

import (
	"fmt"
	"skibidilang/ast"
	"skibidilang/lexer"
	"skibidilang/token"
	"strconv"
)

const (
	_ int = iota
	lowest
	equals      // ==
	lessgreater // > or <
	add         // +
	multiply    // *
	prefix      // -X or !X
	call        // myFunction(X)
)

var precedences = map[token.TokenType]int{
	token.EQ:       equals,
	token.NEQ:      equals,
	token.LT:       lessgreater,
	token.GT:       lessgreater,
	token.ADD:      add,
	token.SUB:      add,
	token.SLASH:    multiply,
	token.ASTERISK: multiply,
}

type Parser struct {
	l              *lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	errors         []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	//Adding prefix functions
	p.prefixParseFns[token.TRUE] = prefixParseFn(p.parseBoolean)
	p.prefixParseFns[token.FALSE] = prefixParseFn(p.parseBoolean)
	p.prefixParseFns[token.IDENT] = prefixParseFn(p.parseIdentifier)
	p.prefixParseFns[token.INT] = prefixParseFn(p.parseInteger)
	p.prefixParseFns[token.NOT] = prefixParseFn(p.parsePrefixExpression)
	p.prefixParseFns[token.SUB] = prefixParseFn(p.parsePrefixExpression)
	p.prefixParseFns[token.LPAREN] = prefixParseFn(p.parseGroupedExpression)
	p.infixParseFns[token.ADD] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.SUB] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.NOT] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.ASTERISK] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.SLASH] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.AND] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.OR] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.LT] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.GT] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.POWER] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.EQ] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.NEQ] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.INC] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.LEQ] = infixParseFn(p.parseInfixExpression)
	p.infixParseFns[token.GEQ] = infixParseFn(p.parseInfixExpression)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.curToken}
	statement.Expression = p.parseExpression(lowest)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseInteger() ast.Expression {
	literal := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	literal.Value = value
	return literal
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(prefix)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(lowest)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.addError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) addError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return lowest
}
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return lowest
}
