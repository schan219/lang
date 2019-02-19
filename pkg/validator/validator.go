package validator

import (
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
