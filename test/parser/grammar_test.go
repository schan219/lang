package parser

import (
	"fmt"
	"reflect"
	"testing"

	"lang/pkg/parser"

	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
)

type TestVal struct {
	Field  string
	ValStr string
	ValInt int32
}

// Only useful for testing function declarations with a string as the body.
type TestFnVal struct {
	Name string
	Args []string
	Body string
}

func TestMainPrimitives(t *testing.T) {
	// Build the assertor and the tokenizer
	tokenizer, err := participle.Build(&parser.Program{})
	assert := assert.New(t)

	if err != nil {
		panic(err)
	}

	// Holds in the values for later
	// Each program is in the form
	// @program:<@field,@value>
	programs := map[string]TestVal{
		`(main () true)`: {
			Field:  "Atom",
			ValStr: "true",
		},
		`(main () b101101)`: {
			Field:  "Atom",
			ValStr: "b101101",
		},
		`(main () 1)`: {
			Field:  "Int",
			ValInt: 1,
		},
		`(main () 3)`: {
			Field:  "Int",
			ValInt: 3,
		},
		`(main () 1000033)`: {
			Field:  "Int",
			ValInt: 1000033,
		},
		`(main () "1")`: {
			Field:  "Str",
			ValStr: "1",
		},
		`(main () "JOE_IS_CoOL!!1")`: {
			Field:  "Str",
			ValStr: "JOE_IS_CoOL!!1",
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
			intOutput := reflect.ValueOf(*root.Main.Body).FieldByName(fieldName).Int()
			assert.Equalf(
				output.ValInt,
				int32(intOutput), // This conversion is because reflect outputs to 64 bit ints
				"Hmm, we failed the int output, for: %s", sourceCode,
			)
		} else {
			assert.Equalf(
				output.ValStr, parsedOutput,
				"Hmm, we failed the string output, for: %s", sourceCode,
			)
		}
	}
}

func TestFnDecl(t *testing.T) {
	// Build the assertor and the tokenizer
	tokenizer, err := participle.Build(&parser.Program{})
	assert := assert.New(t)

	if err != nil {
		panic(err)
	}

	// Holds in the values for later
	// Each program is in the form
	// @program:<@code,@fields>
	programs := make(map[string]TestFnVal)

	code1 := `(defun test (a b) "abc")`

	// Holds in the values for later
	// Each field is in the form
	// @field:<@name,@value>
	// Tests the different fields of only one program.
	programs[code1] = TestFnVal{Name: "test", Args: []string{"a", "b"}, Body: "abc"}

	fmt.Println(programs)

	for sourceCode, output := range programs {
		root := &parser.Program{}
		tokenizer.ParseString(sourceCode, root)
		// Test the non nilness of FnDecl.
		assert.NotNilf(root.Definitions[0].FnDecl, "Hmm, FnDecl is nil for `%s`", sourceCode)
		// Test function name
		parsedOutput := reflect.ValueOf(*root.Definitions[0].FnDecl).FieldByName("Name").String()
		assert.Equalf(output.Name, parsedOutput, "Hmm, we failed the function declaration name, for: %s", sourceCode)
		// Test length of args.
		assert.Equal(len(root.Definitions[0].FnDecl.Args), len(output.Args), "Hmm, the number of args for FnDecl is incorrect.`%s`", sourceCode)
		// Test Args
		for i, value := range output.Args {
			parsedOutput = reflect.ValueOf(*root.Definitions[0].FnDecl.Args[i]).FieldByName("Atom").String()
			assert.Equalf(value, parsedOutput, "Hmm, we failed the function declaration args, for: %s", sourceCode)
		}
		// Test Body
		parsedOutput = reflect.ValueOf(*root.Definitions[0].FnDecl.Body).FieldByName("Str").String()
		assert.Equalf(output.Body, parsedOutput, "Hmm, we failed the function declaration body, for: %s", sourceCode)
	}
}

/*func TestFnCall(t *testing.T) {
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
}*/
