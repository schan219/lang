package main

import (
	"fmt"
	"lang/pkg/parser"
	"github.com/alecthomas/participle"
)

func main () {
	tokenizer, _ := participle.Build(&parser.Program{})

	root := &parser.Program{}
	tokenizer.ParseString("(main () 10)", root)
	fmt.Printf("%+v\n", root.Main.Body)
}