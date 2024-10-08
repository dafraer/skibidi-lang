package ast

import (
	"bytes"
	"skibidilang/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token //the first token of the expression
	Expression Expression
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type Boolean struct {
	Token token.Token
	Value bool
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

// Important useless dummy methods that make  structs implement interfaces
func (b *Boolean) expressionNode()                   {}
func (b *Boolean) TokenLiteral() string              { return b.Token.Literal }
func (b *Boolean) String() string                    { return b.Token.Literal }
func (oe *InfixExpression) expressionNode()          {}
func (oe *InfixExpression) TokenLiteral() string     { return oe.Token.Literal }
func (il *IntegerLiteral) expressionNode()           {}
func (il *IntegerLiteral) TokenLiteral() string      { return il.Token.Literal }
func (il *IntegerLiteral) String() string            { return il.Token.Literal }
func (i *Identifier) expressionNode()                {}
func (i *Identifier) TokenLiteral() string           { return i.Token.Literal }
func (rs *ReturnStatement) statementNode()           {}
func (rs *ReturnStatement) TokenLiteral() string     { return rs.Token.Literal }
func (ls *LetStatement) statementNode()              {}
func (ls *LetStatement) TokenLiteral() string        { return ls.Token.Literal }
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (pe *PrefixExpression) expressionNode()         {}
func (pe *PrefixExpression) TokenLiteral() string    { return pe.Token.Literal }

// String methods
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string { return i.Value }
