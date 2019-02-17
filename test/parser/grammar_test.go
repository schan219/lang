package parser

import (
	_"fmt"
	"testing"
	"reflect"

	"lang/pkg/parser"
	
	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
)

type TestVal struct {
	Field   string
	ValStr  string
	ValInt  int32
}

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
	programs  := map[string]TestVal {
		"(main () true)": TestVal {
			Field:  "Atom",
			ValStr: "true",
		},
		"(main () b101101)": TestVal {
			Field:  "Atom",
			ValStr: "b101101",
		},
		"(main () 1)": TestVal {
			Field:  "Int",
			ValInt: 1,
		},
		"(main () 3)": TestVal {
			Field:  "Int",
			ValInt: 3,
		},
		"(main () 1000033)": TestVal {
			Field:  "Int",
			ValInt: 1000033,
		},
		`(main () "1")`: TestVal {
			Field:  "Str",
			ValStr: "1",
		},
		`(main () "JOE_IS_COOL")`: TestVal {
			Field:  "Str",
			ValStr: "JOE_IS_COOL",
		},
	}


    for sourceCode, output := range programs {
		fieldName := output.Field

		root := &parser.Program{}
		tokenizer.ParseString(sourceCode, root)

		// Test the non nillness of main.
		assert.NotNilf(root.Main, "Hmm, main is nil for `%s`", sourceCode)

		// Test if parsed value is expected.
		parsedOutput := reflect.ValueOf(*root.Main.Body).FieldByName(fieldName).String()

		// We must treat ints different to strings
		if parsedOutput == "<int32 Value>" {
			intOuput := reflect.ValueOf(*root.Main.Body).FieldByName(fieldName).Int()
			assert.Equalf(
				output.ValInt,
				int32(intOuput),	// This conversion is because reflect outputs to 64 bit ints
				"Hmm, we faild the int output, for: %s", sourceCode,
			)
		} else {
			assert.Equalf(
				output.ValStr, parsedOutput,
				"Hmm, we failed the string output, for: %s", sourceCode,
			)
		}
    }
}

/** I will refactor these later -- Joe

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
*/