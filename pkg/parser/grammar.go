package parser

type Program struct {
	Defs []*Def `"(" @@ ")"`
}

type Def struct {
	Function  *FnDecl   `"defun" @@`
	Variable *VarDecl   `| "defvar" @@`
	Output *OutDecl    `| "defoutput" @@`
}

type OutDecl struct {
	
	Number *int `@Int`
	Name *string  `@Ident`
	Args []*Expr `"(" (@@)* ")"`
	Script *Expr   `@@`
}

type FnDecl struct {
	Name string  `@Ident`
	FnPtr uint32
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
	VarPtr uint32 
	Str  string  `| @String`
	Int  int32   `| @Int`
}
