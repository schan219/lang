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
	Loop         *Loop         `| @@`
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

type Loop struct {
	Start *Expr `"(" "for" "(" "(" @@ ")"`
	End   *Expr `"(" "<" "i" @@ ")"`
	Inc   *Expr `"(" "+" "i" @@ ")" ")"`
	Body  *Expr `@@ ")"`
}

type FunctionCall struct {
	Name      string  `"(" @Ident`
	Arguments []*Expr `(@@)* ")"`
}
