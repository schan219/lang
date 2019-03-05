package parser

type Program struct {
	DefOrMain []*DefOrMain `"(" @@ ")"`
}

type DefOrMain struct {
	FnDecl  *FnDecl  `"defun" @@`
	VarDecl *VarDecl `| "defvar" @@`
	Main    *Main    `| "main" @@`
}

type Main struct {
	Args []*Expr `"(" (@@)* ")"`
	Body *Expr   `@@`
}

type FnDecl struct {
	Name string  `@Ident`
	Args []*Expr `"(" (@@)* ")"`
	Body *Expr   `@@`
}

type VarDecl struct {
	Name  string `@Ident`
	Value *Expr  `@@`
}

type FnCall struct {
	Name string  `@Ident`
	Args []*Expr `(@@)*`
}

type Expr struct {
	Fn   *FnCall `"(" @@ ")"`
	Atom string  `| @Ident`
	Str  string  `| @String`
	Int  int32   `| @Int`
}
