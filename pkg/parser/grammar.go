package parser

type Program struct {
	Defs []*Def `"(" @@ ")"`
}

type Def struct {
	Function  *FnDecl   `"defun" @@`
	Variable *VarDecl   `| "defvar" @@`
	Output  *OutDecl    `| "defoutput" @@`
}

type OutDecl struct {
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
