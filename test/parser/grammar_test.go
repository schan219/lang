package parser

import (
	"../../pkg/parser"
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestMainPrimitives (t *testing.T) {
	// Build the assertor and the tokenizer
	tokenizer, err := participle.Build(&parser.Program{})
	assert := assert.New(t)

	if err != nil {
		panic(err)
	}

	// Holds in the values for later
	// Each program is in the form
	// @program:<@field,@value> 
	//
	programs  := map[[2]string]{
		"(main () true)": []string {
			"Atom",
			"true"
		},
		"(main () b101101)": []string {
			"Atom",
			"b101101"
		},
		
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
