// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package scanner is a thin wrapper of Go scanner package.
// Rat scanner returns Rat token instead of Go token.
// It's the only difference between these two packages.
package scanner

import (
	goscanner "go/scanner"
	gotoken "go/token"

	"github.com/go-rat/rat/token"
)

type Scanner struct {
	goscanner.Scanner
}

// Scan scans the next token and returns the token position, the token,
// and its literal string if applicable.
// The source end is indicated by token.EOF.
//
// Scan works exactly the same as Go Scanner.Scan.
// More details see Go package document.
// http://golang.org/pkg/go/scanner/#Scanner.Scan
func (scanner *Scanner) Scan() (pos gotoken.Pos, tok token.Token, lit string) {
	pos, gotok, lit := scanner.Scanner.Scan()

	tok, lit = token.Convert(gotok, lit)
	return pos, tok, lit
}
