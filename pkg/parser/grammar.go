package parser

type Program struct {
	DefOrMain *DefOrMain `@@`
}

type DefOrMain struct {
	Definition *Definition `@@`
	Main       *Main       `| @@`
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
	Function  *Function `@@`
	Boolean   string    `| @"true" | @"false"`
	BinString string    `| "b"@Int`
	Num       float64   `| @Float | @Int`
	Name      string    `| @Ident`
	String    string    `| @String`
}

type Function struct {
	Name string  `"(" @Ident`
	Args []*Expr `(@@)* ")"`
}
