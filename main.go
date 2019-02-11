package main

import (
	"github.com/alecthomas/participle"
	"fmt"
)
type Node struct {
	Function   string `"(" (@Ident|"+")`
	Arguments []*Argument `(@@)* ")"`
}

type Argument struct {
	Arg		  []*Node `@@`
	String    string  `| @String`
	Num       float64 `| @Float | @Int`
}

func main () {
	parser, err := participle.Build(&Node{})
	
	if err != nil {
		panic(err)
	}

	root := &Node{}

	fmt.Printf("%+v\n", root)
	parser.ParseString(`(joe "hi" 3  (+ 10 10))`, root)

	fmt.Printf("%+v\n", root)
}
