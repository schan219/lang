package main

import (
	"fmt"
	"lang/internals/cli"
	"lang/pkg/parser"
	"github.com/alecthomas/participle"
)

func main() {
	conf := cli.Start()

	// TODO: manage dependencies here.

	// Tokenize the result.
	ASTBuilder, err := participle.Build(&parser.Program{})
	
	if err != nil {
		panic(fmt.Sprintf("Issue with building parser: %+v", err))
	}

	program := parser.Program{}
	err = ASTBuilder.ParseString(string(conf.Contents), &program)

	if err != nil {
		panic(fmt.Sprintf("Issue parsing the file... %+v", err))
	}
	
	fmt.Println("%+v", program);
	// Translate each token starting with the definitions.
	// Clean up / Execute tests.
}
