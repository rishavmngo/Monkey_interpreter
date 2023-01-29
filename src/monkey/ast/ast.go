package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// program have a Statements array which contains memember which is Statement type
type Program struct {
	Statements []Statement
}

// this will give the TokenLiteral of the first statement from all the Statements
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//LetStatement
// let statement have three parts first is let token then the name of variable then third is value/expression
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

//satisfy Statment interface
func (ls *LetStatement) statementNode()       {}
//satisfy Node interface
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier type
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Return Statement
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

