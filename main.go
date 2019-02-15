package main

import (
	"fmt"
	"github.com/alecthomas/participle"
)

type Program struct {
	DefOrMain *DefOrMain `@@`
}

type DefOrMain struct {
	Definition *Definition `@@`
	Main       *Main       `| @@`
}

type Definition struct {
	FunctionDecl *FunctionDecl `@@`
}

type Main struct {
	Args []*Expr `"(" "main" "(" (@@)* ")" `
	Body *Expr   `@@ ")"`
}

type FunctionDecl struct {
	Name string  `"(" "define-fn" @Ident `
	Args []*Expr `"(" (@@)* ")" `
	Body *Expr   `@@ ")"`
}

type Expr struct {
	Branch       *Branch       `@@`
	Cond         *Cond         `| @@`
	LogicalOp    []*Expr       `| "(" ("and"|"or") @@ (@@)+ ")"`
	FunctionCall *FunctionCall `| @@`
	Boolean      string        `| @"true" | @"false"`
	BinString    string        `| "b"@Int`
	Num          float64       `| @Float | @Int`
	String       string        `| @String`
}

type Branch struct {
	Condition *Expr `"(" "if" @@`
	Then      *Expr `@@`
	Else      *Expr `@@ ")"`
}

type Cond struct {
	Cond  string  `"(" "cond" `
	Cases []*Expr `("[" @@ @@ "]")+`
	Else  []*Expr `("[" "else" @@ "]")? ")"`
}

type FunctionCall struct {
	Name      string  `"(" @Ident`
	Arguments []*Expr `(@@)* ")"`
}

func main() {
	parser, err := participle.Build(&Program{})

	if err != nil {
		panic(err)
	}

	root := &Program{}
	parser.ParseString(`(main () (joe "hi" (jo 10 10) 3 b10))`, root)

	fmt.Printf("%+v\n", root.DefOrMain.Main.Body.FunctionCall.Arguments[0].String)
}
