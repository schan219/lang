package parser

import (
	"fmt"
	"github.com/alecthomas/participle"
)

type Parser struct {
	FnMap map[int]*FnDecl
	FnPtrMap map[string]int

	VarMap map[int]*VarDecl
	VarPtrMap map[string]int

	OutputList []*OutDecl
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(program string) []Defs, error {
	var root Program{}
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		return nil, error(fmt.Sprintf("Tokenizing failed: %+v", err))
	}

	tokenizer.ParseString(program, &root)

	if p.updateMap(&root) != nil {
		return nil, error("Could not view error")
	}
}


func (p *Parser) updateMap(root &Program) error {
	for def,_ := Program.Defs {
		if def.Output != nil {

		} else if def.Function != nil {
			id := len(p.FnMap)
			p.VarMap[id] = def.Function
			p.VarPtrMap[def.Variable.Name] = id
			
		} else if def.Variable != nil {
			id := len(p.VarMap)
			p.VarMap[id] = def.Variable
			p.VarPtrMap[def.Variable.Name] = id
		} else {
			return error(fmt.Sprintf("Could not format this part of the program.."))
		}
	}

	return nil
}