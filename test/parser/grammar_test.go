package parser

import (
	"../../pkg/parser"
	"fmt"
	"github.com/alecthomas/participle"
	"testing"
)

func TestMain(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	tokenizer.ParseString(`(main () (joe "hi" (jo 10 10) 3 b10))`, root)
	v := root.DefOrMain.Main.Body.Function.Args[0].String

	if v != "hi" {
		t.Error("Expected \"hi\", got ", v)
	}
}

func TestBoolean(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	tokenizer.ParseString(`(main () true)`, root)
	v := root.DefOrMain.Main.Body.Boolean

	if v != "true" {
		t.Error("Expected \"true\", got ", v)
	}
}

func TestBinString(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	tokenizer.ParseString(`(main () "b123")`, root)
	v := root.DefOrMain.Main.Body
	fmt.Printf("%+v\n", v)

	if v.BinString != "b123" {
		t.Error("Expected \"b123\", got ", v.BinString)
	}
}

func TestNum(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	tokenizer.ParseString(`(main () 1)`, root)
	v := root.DefOrMain.Main.Body.Num

	if v != 1 {
		t.Error("Expected \"1\", got ", v)
	}
}

func TestSymbol(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	str := `(main () ALICE_ADDR)`

	tokenizer.ParseString(str, root)

	v := root.DefOrMain.Main.Body

	if v.Name != "ALICE_ADDR" {
		t.Error("Expected \"ALICE_ADDR\" got ", v.Name)
	}
}

func TestString(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	str := `(main () "a")`

	tokenizer.ParseString(str, root)

	v := root.DefOrMain.Main.Body

	if v.String != "a" {
		t.Error("Expected \"a\", got ", v.String)
	}
}

func TestFunctionDecl(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}

	str := `(defun test (a)
				(if (= a 0)
					true
					false))`

	tokenizer.ParseString(str, root)

	v := root.DefOrMain.Definition.FunctionDecl

	if v.Name != "test" {
		t.Error("Expected \"test\", got ", v.Name)
	}

	if len(v.Args) != 1 {
		t.Error("Expected \"1\", got ", len(v.Args))
	}

	if v.Args[0].Name != "a" {
		t.Error("Expected \"a\", got ", v.Args[0].Name)
	}

	if v.Body.Function.Name != "if" {
		t.Error("Expected \"if\" got ", v.Body.Function.Name)
	}
}

func TestFunction(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	str := `(main () (eat 1 2 3))`
	tokenizer.ParseString(str, root)

	v := root.DefOrMain.Main

	if len(v.Args) != 0 {
		t.Error("Expected \"0\", got ", len(v.Args))
	}

	if v.Body.Function.Name != "eat" {
		t.Error("Expected \"eat\" got ", v.Body.Function.Name)
	}

	if len(v.Body.Function.Args) != 3 {
		t.Error("Expected \"3\", got ", len(v.Body.Function.Args))
	}

	if v.Body.Function.Args[0].Num != 1 {
		t.Error("Expected \"1\", got ", v.Body.Function.Args[0].Num)
	}

	if v.Body.Function.Args[1].Num != 2 {
		t.Error("Expected \"2\", got ", v.Body.Function.Args[1].Num)
	}

	if v.Body.Function.Args[2].Num != 3 {
		t.Error("Expected \"3\", got ", v.Body.Function.Args[2].Num)
	}
}
