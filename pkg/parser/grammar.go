package parser

type Program struct {
	Defs []*Defs `"(" @@ ")"`
}

type Defs struct {
	FnDecl  *FnDecl  `"defun" @@`
	VarDecl *VarDecl `| "defvar" @@`
	Output  *Output    `| "defoutput" @@`
	Input  *Input    `| "input" @@`
}

type Output struct {
	Name *string  `@Ident`
	Args []*string `(@Ident)*`
	Script *Expr   `@@`
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
