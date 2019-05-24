package main

import (
	"os"
	"flag"
	"lang/internals/help"
	_"lang/internals/cli"
)

const (
	HELP_COMMAND = "help"
	COMPILE_COMMAND = "compile"
)

var (
	output *string = flag.String("-o", "", "Output TX file")
	input string = ""
)

func main() {
	flag.Parse()
//	var conf cli.InitConfig

	if len(os.Args) == 1 {
		panic("No command or input files. Terminated.")
	}
	
	switch os.Args[1] {
		case HELP_COMMAND:
			if len(os.Args) < 3 {
				panic("Cannot help, no command specified")
			} else {
				help.DescribeCommand(os.Args[2])
			}
		case COMPILE_COMMAND:
		//	conf = cli.Build(os.Args[1], output)
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


