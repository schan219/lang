package validator

import (
	"fmt"

	"lang/pkg/parser"
)

type Def struct {
	DefName   string
	DefType   string
	DefIndex  int
}

// Stores the definition locations
var DefMap map[string]Def

func Validate (prgm parser.Program) {
	// TODO: validate that there's only one body in Main
	
	for i,def := range prgm.Definitions {
		switch {
		case def.VarDecl != nil:
			name := def.VarDecl.Name

			if _,ok := DefMap[name]; ok {
				panic(fmt.Sprintf("%s declared redeclared",name))
			}
			
			// Since we passed all barriers we add it to the map
			DefMap[name] = Def{
				DefName: name,
				DefType: "variable",
				DefIndex: i,
			}
		case def.FunctionDecl != nil:
			name := def.FunctionDecl.Name

			if _,ok := DefMap[name]; ok {
				panic(fmt.Sprintf("%s declared redeclared",name))
			}

			// Since we passed all barriers we add it to the map
			DefMap[name] = Def{
				DefName: name,
				DefType: "function",
				DefIndex: i,
			}
		}
	}
}
