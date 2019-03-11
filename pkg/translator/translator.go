package translator

import (
	"lang/pkg/"
	"lang/pkg/parser"
	"lang/pkg/script"
)

func Translate(node *parser.Expr) []code {

}

func translateCond(code []byte, node *parser.FnCall) []byte {
	// The actual operation must be caller validated.
	// If it's a conditional must have only 2 arguments.
	if len(node.Args) != 2 {
		panic(fmt.Sprintf(
			"Contains either too many or too little args, should be 2. \n %s",
			"(operation arg1 arg2)\n",
		))
	}

	var baseOp int := 0;

	switch node.Name {
	case ">":
		baseOp = script.OP_GREATERTHAN
	case "<":
		baseOp = script.OP_LESSTHAN
	case "=":
		baseOp = script.OP_NUMEQUAL
	case "<=":
		baseOp = script.OP_LESSTHANOREQUAL
	case ">=":
		baseOp = script.OP_LESSTHANOREQUAL
	case "!=":
		baseOp = script.OP_LESSTHANOREQUAL
	default:
		panic(fmt.Sprintf(
			"Uncaught conditional please report this! %s",
			node.Name,
		))
	}

	// Stack should be [Data1] [Data2] [Operation]
	code = append(
		code,
		Translate(node.Args[0]),
		Translate(node.Args[1]),
		baseOp,
	)

	return code
}

func translateIF(code []byte, node *parser.FnCall) []byte {
	// Assert the validity of IF.
	if len(node.Args) != 3 {
		panic(fmt.Sprintf(
			"If does not contain enough arguments should be in the form\n %s",
			"(if (...conditional) (...) (...))\n",
		))
	}

	// Add in the IF statement.
	code = append(
		code,
		Translate(node.Args[0]),
		script.OP_IF,
		Translate(node.Args[1]),
		script.OP_ELSE,
		Translate(node.Args[2]),
		script.OP_ENDIF,
	)

	return code
}
