package main

import (
	"fmt"
	"lang/internals/cli"
	"lang/pkg/parser"
	"github.com/alecthomas/participle"
)

func main() {
	// Grab the input from somewhere
	cli.Start()

	//
	// We should manage dependencies here.
	//

	// Tokenize the result.
	_, err := participle.Build(&parser.Program{})
	
	if err != nil {
		panic(fmt.Sprintf("Issue with building parser: %+v", err))
	}

	// Translate each token starting with the definitions.
	// Clean up / Execute tests.
}
