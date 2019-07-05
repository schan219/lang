// This is the package that both lexes and parses a provided
// string of Lang. Initially this was broken up into 2 packages
// but the simplicity of S-expressions allows us to marry the two
// in a single package.
package parser

import (
	"errors"
)

type SExpr struct {
	// We use pointers because they could be nil
	Str *string
	Num *string

	// If we have other SExpressions nested in this
	Args []*SExpr
}

type Parser struct {
	defs []*SExpr
	raw string
}

func NewParser() *Parser {
	return &Parser{}
}

// Parse takes in a string and updates the parser object with the global scopes defs.
// Will return an error if an error happens.
func (p *Parser) Parse(str string) error {
	p.raw = str

	for p.Trim() {
		if err := p.parseBlock(); err {
			return err
		}
	}

	if len(p.defs) == 0 {
		return errors.New("Could not find bracket '(', \n empty file detected")
	}

	return nil
}

// Trim updates the raw value until it starts with '('.
// Returns false if there is no '(' found.
func (p *Parser) Trim() bool {
	for i,_ := range p.raw {
		if p.raw[i] == '(' {
			p.raw = p.raw[(i - 1):]
			return true
		}
	}

	return false
}
