// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package token extends Go token package to support Rat token.
package token

import (
	gotoken "go/token"
)

type Token gotoken.Token

// All Rat tokens use negative value to avoid conflict with Go tokens.
const (
	GENERIC Token = -iota - 1
	YIELD
)

// All Rat tokens which are the same as Go tokens.
const (
	// Special tokens
	ILLEGAL = Token(gotoken.ILLEGAL)
	EOF     = Token(gotoken.EOF)
	COMMENT = Token(gotoken.COMMENT)

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT  = Token(gotoken.IDENT)
	INT    = Token(gotoken.INT)
	FLOAT  = Token(gotoken.FLOAT)
	IMAG   = Token(gotoken.IMAG)
	CHAR   = Token(gotoken.CHAR)
	STRING = Token(gotoken.STRING)

	// Operators and delimiters
	ADD = Token(gotoken.ADD)
	SUB = Token(gotoken.SUB)
	MUL = Token(gotoken.MUL)
	QUO = Token(gotoken.QUO)
	REM = Token(gotoken.REM)

	AND     = Token(gotoken.AND)
	OR      = Token(gotoken.OR)
	XOR     = Token(gotoken.XOR)
	SHL     = Token(gotoken.SHL)
	SHR     = Token(gotoken.SHR)
	AND_NOT = Token(gotoken.AND_NOT)

	ADD_ASSIGN = Token(gotoken.ADD_ASSIGN)
	SUB_ASSIGN = Token(gotoken.SUB_ASSIGN)
	MUL_ASSIGN = Token(gotoken.MUL_ASSIGN)
	QUO_ASSIGN = Token(gotoken.QUO_ASSIGN)
	REM_ASSIGN = Token(gotoken.REM_ASSIGN)

	AND_ASSIGN     = Token(gotoken.AND_ASSIGN)
	OR_ASSIGN      = Token(gotoken.OR_ASSIGN)
	XOR_ASSIGN     = Token(gotoken.XOR_ASSIGN)
	SHL_ASSIGN     = Token(gotoken.SHL_ASSIGN)
	SHR_ASSIGN     = Token(gotoken.SHR_ASSIGN)
	AND_NOT_ASSIGN = Token(gotoken.AND_NOT_ASSIGN)

	LAND  = Token(gotoken.LAND)
	LOR   = Token(gotoken.LOR)
	ARROW = Token(gotoken.ARROW)
	INC   = Token(gotoken.INC)
	DEC   = Token(gotoken.DEC)

	EQL    = Token(gotoken.EQL)
	LSS    = Token(gotoken.LSS)
	GTR    = Token(gotoken.GTR)
	ASSIGN = Token(gotoken.ASSIGN)
	NOT    = Token(gotoken.NOT)

	NEQ      = Token(gotoken.NEQ)
	LEQ      = Token(gotoken.LEQ)
	GEQ      = Token(gotoken.GEQ)
	DEFINE   = Token(gotoken.DEFINE)
	ELLIPSIS = Token(gotoken.ELLIPSIS)

	LPAREN = Token(gotoken.LPAREN)
	LBRACK = Token(gotoken.LBRACK)
	LBRACE = Token(gotoken.LBRACE)
	COMMA  = Token(gotoken.COMMA)
	PERIOD = Token(gotoken.PERIOD)

	RPAREN    = Token(gotoken.RPAREN)
	RBRACK    = Token(gotoken.RBRACK)
	RBRACE    = Token(gotoken.RBRACE)
	SEMICOLON = Token(gotoken.SEMICOLON)
	COLON     = Token(gotoken.COLON)

	// Keywords
	BREAK    = Token(gotoken.BREAK)
	CASE     = Token(gotoken.CASE)
	CHAN     = Token(gotoken.CHAN)
	CONST    = Token(gotoken.CONST)
	CONTINUE = Token(gotoken.CONTINUE)

	DEFAULT     = Token(gotoken.DEFAULT)
	DEFER       = Token(gotoken.DEFER)
	ELSE        = Token(gotoken.ELSE)
	FALLTHROUGH = Token(gotoken.FALLTHROUGH)
	FOR         = Token(gotoken.FOR)

	FUNC   = Token(gotoken.FUNC)
	GO     = Token(gotoken.GO)
	GOTO   = Token(gotoken.GOTO)
	IF     = Token(gotoken.IF)
	IMPORT = Token(gotoken.IMPORT)

	INTERFACE = Token(gotoken.INTERFACE)
	MAP       = Token(gotoken.MAP)
	PACKAGE   = Token(gotoken.PACKAGE)
	RANGE     = Token(gotoken.RANGE)
	RETURN    = Token(gotoken.RETURN)

	SELECT = Token(gotoken.SELECT)
	STRUCT = Token(gotoken.STRUCT)
	SWITCH = Token(gotoken.SWITCH)
	TYPE   = Token(gotoken.TYPE)
	VAR    = Token(gotoken.VAR)
)

const (
	LowestPrec  = gotoken.LowestPrec // non-operators
	UnaryPrec   = gotoken.UnaryPrec
	HighestPrec = gotoken.HighestPrec
)

const (
	genericIdent = "generic"
	yieldIdent   = "yield"
)

// Looks up an identifier and returns Rat token.
func Lookup(ident string) Token {
	t := gotoken.Lookup(ident)

	if t == gotoken.IDENT {
		if ident == genericIdent {
			return GENERIC
		}

		if ident == yieldIdent {
			return YIELD
		}
	}

	return Token(t)
}

// Converts a Go token to Rat token.
func Convert(tok gotoken.Token, lit string) (Token, string) {
	if tok == gotoken.IDENT {
		if lit == genericIdent {
			return GENERIC, lit
		}

		if lit == yieldIdent {
			return YIELD, lit
		}
	}

	return Token(tok), lit
}

// IsKeyword returns true for keyword tokens and false otherwise.
func (t Token) IsKeyword() bool {
	if t == GENERIC || t == YIELD {
		return true
	}

	return gotoken.Token(t).IsKeyword()
}

// IsLiteral returns true for identifiers and basic type literals.
// It returns false otherwise.
func (t Token) IsLiteral() bool {
	return gotoken.Token(t).IsLiteral()
}

// IsOperator returns true for operators and delimiters.
// It returns false otherwise.
func (t Token) IsOperator() bool {
	return gotoken.Token(t).IsOperator()
}

// Precendence returns the operator precedence of the binary operator op.
// If op is not a binary operator, the result is LowestPrec.
func (t Token) Precedence() int {
	return gotoken.Token(t).Precedence()
}

// String returns the string for the token.
// String of operators, delimiters and keywords are the actual token
// character sequence. For example, string of GENERIC is "generic",
// and string of ADD is "+".
// For all other tokens, String returns token constant name like "IDENT".
func (t Token) String() string {
	if t == GENERIC {
		return genericIdent
	}

	if t == YIELD {
		return yieldIdent
	}

	return gotoken.Token(t).String()
}
