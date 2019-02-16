package parser

import (
	_"fmt"
	"testing"
	"reflect"

	"lang/pkg/parser"
	
	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
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
	programs  := map[string][2]string {
		"(main () true)": [2]string {
			"Atom",
			"true",
		},
		"(main () b101101)": [2]string {
			"Atom",
			"b101101",
		},
		"(main () 1)": [2]string {
			"Int",
			"10",
		},
		"(main () 3)": [2]string {
			"Int",
			"10",
		},
		"(main () 1000033)": [2]string {
			"Int",
			"1000033",
		},
		`(main () "JOE_IS_COOL")`: [2]string {
			"Str",
			"JOE_IS_COOL",
		},
	}


    for sourceCode, output := range programs {
		fieldName := output[0]
		fieldVal := output[1]

		root := &parser.Program{}
		tokenizer.ParseString(sourceCode, root)

		// Test the non nillness of main.
		assert.NotNilf(root.Main, "Hmm, main is nil for `%s`", sourceCode)

		// Test if parsed value is expected.
		parsedOutput := reflect.ValueOf(*root.Main).FieldByName(fieldName).String()
		assert.Equalf(
			parsedOutput, fieldVal,
			"Hmm, we failed, for: %s", sourceCode,
		)
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