package parser

import (
	"fmt"
	"github.com/alecthomas/participle"
)

type Parser struct {
	tokenizer 
	FnMap map[int]*Def
	VarMap map[int]*Def
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(program string) []Defs {
	var root Program{}
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(fmt.Sprintf("Tokenizing failed: %+v", err))
	}

	tokenizer.ParseString(program, &root)
	p.updateMap(&root)
}

func (p *Parser) updateMap(root &Program) {
	for def,_ := Program.Defs {
		switch def {
		case def.FnDecl != nil:
			case def.VarDecl != nil:
				case  
		}
	}
}