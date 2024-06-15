package parser

import (
	"skibidilang/ast"
	"skibidilang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
skibidi x = 5;
skibidi y = 10;
skibidi foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []string{"x", "y", "foobar"}
	for i, expectedIdentifier := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, expectedIdentifier) {
			return
		}
	}
}
func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "skibidi" {
		t.Errorf("s.TokenLiteral not 'skibidi'. got=%q", statement.TokenLiteral())
		return false
	}
	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", statement)
		return false
	}
	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}
	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("statement.Name not '%s'. got=%s", name, letStatement.Name)
		return false
	}
	return true
}
func TestReturnStatements(t *testing.T) {
	input := `
	goon 5;
	goon 10;
	goon 993322;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not *ast.returnStatement. got=%T", statement)
			continue
		}
		if returnStatement.TokenLiteral() != "goon" {
			t.Errorf("returnStatement.TokenLiteral not 'goon', got %q",
				returnStatement.TokenLiteral())
		}
	}
}
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
