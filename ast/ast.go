package ast

import "monkey-lang/token"

// The base Node interface
type Node interface {
	TokenLiteral() string
}

// All statement nodes implement this
type Statement interface {
	Node
	stateMentNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) >= 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.Let
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) stateMentNode() {

}

// Expression
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
}

type ReturnStatement struct {
	Token       token.Token //token.RETURN
	ReturnValue Expression
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) stateMentNode() {
}
