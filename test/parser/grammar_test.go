package parser

import (
	"fmt"
	"github.com/alecthomas/participle"
	"../../pkg/parser"
	"testing"
)

func TestMain(t *testing.T) {
	tokenizer, err := participle.Build(&parser.Program{})

	if err != nil {
		panic(err)
	}

	root := &parser.Program{}
	tokenizer.ParseString(`(main () (joe "hi" (jo 10 10) 3 b10))`, root)
	v := root.DefOrMain.Main.Body.FunctionCall.Arguments[0].String
	fmt.Printf("%+v\n", v)

	if (v != "hi") {
		t.Error("Expected \"hi\", got ", v)
	}
}
