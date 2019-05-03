package parser

type Program struct {
	Defs []*Defs `"(" @@ ")"`
}

type Defs struct {
	FnDecl  *FnDecl  `"defun" @@`
	VarDecl *VarDecl `| "defvar" @@`
	Output  *Output    `| "output" @@`
	Input  *Input    `| "input" @@`
}

type Output struct {
	OutputIndex *int  `@Int`
	Value  *float64 `@Float`
	ScriptPubKey *Expr   `@@`
}

type Input struct {
	InputIndex *int  `@Int`
	InputHash  *string `@Ident`
	InputHashIndex *int `@Int`
	ScriptSig *Expr   `@@`
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
