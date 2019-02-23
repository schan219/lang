package translator

import (
	"lang/pkg/script"
	"lang/pkg/parser"
	"lang/pkg/"
)

func Translate(node *parser.Expr)[]code {

}


func translateIF(code []byte, node *parser.FnCall) []byte {
	// Assert the validity of IF.
	if len(node.Args) != 3 {
		panic(fmt.Sprintf(
			"If does not contain enough arguments should be in the form\n %s",
			"(if (...conditional) (...) (...))\n"
		))
	}

	// Add in the IF statement.
	code = append(
		code,
		Translate(node.Args[0])
		script.OP_IF,
		Translate(node.Args[1])
		script.OP_ELSE,
		Translate(node.Args[2])
		script.OP_ENDIF,
	)

	return code
}










