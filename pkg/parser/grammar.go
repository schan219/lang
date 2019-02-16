package parser

type Program struct {
	Definitions  []*Definition `@@`
	Main         *Main       `| @@`
}

type Definition struct {
	FunctionDecl *FunctionDecl `@@`
	VarDecl      *VarDecl      `@@`
}

type Main struct {
	Args []*Expr `"(" "main" "(" (@@)* ")" `
	Body *Expr   `@@ ")"`
}

type FunctionDecl struct {
	Name string  `"(" "defun" @Ident `
	Args []*Expr `"(" (@@)* ")" `
	Body *Expr   `@@ ")"`
}

type VarDecl struct {
	Name  string `"(" "defvar" @Ident `
	Value *Expr  `@@ ")"`
}

type Expr struct {
	Atom  string  `@Ident`
	Str   string  `| @String`
	Val   int32   `| @Int`
	Float float64 `| @Float`
}