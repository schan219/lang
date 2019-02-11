package main

import (
	"github.com/alecthomas/participle"
	"fmt"
)
type Node struct {
	Function   string `"("@Ident`
	Arguments []*Argument `@@ ")"`
}

type Argument struct {
	String    string  `@String`
	Num       float64     `| @Float`
}

func main () {
	parser, err := participle.Build(&Node{})
	
	if err != nil {
		panic(err)
	}

	root := &Node{}

	fmt.Printf("%+v\n", root)
	parser.ParseString(`(joe 10.2)`, root)

	fmt.Printf("%+v\n", root)
}
