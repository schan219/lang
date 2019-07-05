package main

import (
	"os"
	"flag"
	"fmt"
	"lang/internals/help"
	_"lang/internals/cli"
)

const (
	DOC_COMMAND = "doc"
	COMPILE_COMMAND = "compile"
	LANG_COMMAND = "lang"
)

var (
	output *string = flag.String("-o", "", "Output TX file")
	input string = ""
)

func main() {
	flag.Parse()
//	var conf cli.InitConfig

	if len(os.Args) == 1 {
		fmt.Println("No command or input files. Terminated.")
	}
	
	switch os.Args[1] {
		case DOC_COMMAND:
			if len(os.Args) < 3 {
				fmt.Printf("doc works via: %s doc <function>\n", LANG_COMMAND)
			} else {
				help.DescribeCommand(os.Args[2])
			}
		case COMPILE_COMMAND:
			conf = cli.Build(os.Args[1], output)
			
	}
}




/**



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
*/


