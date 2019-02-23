package parser

type Program struct {
	Definitions []*Definition `@@`
	Main        *Main         `| @@`
}

type Definition struct {
	FnDecl  *FnDecl  `@@`
	VarDecl *VarDecl `@@`
}

type Main struct {
	Args []*Expr `"(" "main" "(" (@@)* ")" `
	Body *Expr   `@@ ")"`
}

type FnDecl struct {
	Name string  `"(" "defun" @Ident `
	Args []*Expr `"(" (@@)* ")" `
	Body *Expr   `@@ ")"`
}

type VarDecl struct {
	Name  string `"(" "defvar" @Ident `
	Value *Expr  `@@ ")"`
}

type FnCall struct {
	Name string  `"(" @Ident`
	Args []*Expr `(@@)* ")"`
}

type Expr struct {
	Fn   *FnCall `@@`
	Atom string  `| @Ident`
	Str  string  `| @String`
	Int  int32   `| @Int`
}
