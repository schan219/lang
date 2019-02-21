package validator

import (
<<<<<<< HEAD
	"fmt"

	"lang/pkg/parser"
)

type Def struct {
	DefName   string
	DefType   string
	DefIndex  int
	ArgCount  int
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
				ArgCount: len(def.FunctionDecl.Args)
			}
		}
	}

	// Statically analyze each function
	for i,def := range prgm.Definitions {
		validateFn(def)
	}
}

func validateFn (def parser.FunctionDecl) {
	if (def.Body) {
		
	}
}

=======
	"lang/pkg/parser"
)


func Validate(program parser.Program) {

	// Program must have a main function
	if program.Main == nil {
		panic("Please insert a main function into your program\n"
		+ "i.e. (main () true)")
	}

	for def,_ := range program.Definitions {
		validateFn(def.Name, def.Args, def.Expr)
	}
}

func validateFn(name string, args parser.Expr, body parser.Expr) {

	// There should only be T
}
>>>>>>> joe/interpreter
