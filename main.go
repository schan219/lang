package main

import (
	"fmt"
	"flag"
	"os"

	"lang/internals/cli"
	"lang/pkg/parser"
	"github.com/alecthomas/participle"
)


var (
	output string = flag.String("-o", "", "Output TX file")
	input string = ""
)


func main() {
	flag.Parse()
	var conf cli.InitConfig

	if len(os.Args) == 1 {
		conf = cli.Start()
	} else {
		conf = cli.Build(os.Args[1], output) 
	}

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
