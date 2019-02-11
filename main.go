package main

import (
	"github.com/alecthomas/participle"
	"fmt"
)
type Node struct {
	Function   string `"("@Ident")"`
}

type Argument struct {
	String    string  `@String`
	Int       int     `@Float`
}

func main () {
	parser, err := participle.Build(&Node{})
	
	if err != nil {
		panic(err)
	}

	root := &Node{}

	fmt.Printf("%+v\n", root)
	parser.ParseString(`(joe)`, root)

	fmt.Printf("%+v\n", root)
}
