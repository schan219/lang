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
	fmt.Printf("%+v\n", v)

	if v != "hi" {
		t.Error("Expected \"hi\", got ", v)
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
	fmt.Printf("%+v\n", v.Body)

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
