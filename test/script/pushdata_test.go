package script

import (
	"fmt"
	"testing"

	"lang/pkg/script"
	"github.com/stretchr/testify/assert"
)

type Frame script.Frame;
type Stack script.Stack;

type TestVal struct {
	FinalStack   [][]byte
	IsError      bool
	Truthy		 bool
}

func TestMainPrimitives (t *testing.T) {
	assert := assert.New(t)
	err, intp := script.NewInterpreter("ALL");

	if err != nil {
		panic(err)
	}

	programs := map[string]TestVal {
		string([]byte{script.OP_0}): TestVal {
			FinalStack: [][]byte{[]byte{}},
			IsError: false,
			Truthy: false,
		},
		string([]byte{script.OP_DATA_1,script.OP_DATA_2}): TestVal {
			FinalStack: [][]byte{[]byte{script.OP_DATA_2}},
			IsError: false,
			Truthy: true,
		},
		string([]byte{script.OP_DATA_5,script.OP_DATA_2,script.OP_DATA_3,script.OP_DATA_5,script.OP_DATA_7,script.OP_DATA_6}): TestVal {
			FinalStack: [][]byte{[]byte{
				script.OP_DATA_2,script.OP_DATA_3,script.OP_DATA_5,script.OP_DATA_7,script.OP_DATA_6,
			}},
			IsError: false,
			Truthy: true,
		},
	}

	for sourceCode, output := range programs {
		err, truthy := intp.Exec([]byte(sourceCode));
		fmt.Printf("%+v\n", intp.Stack);

		if output.IsError {
			assert.NotNil(err);
		} else {
			assert.Nil(err);
		}

		assert.Equal(intp.Stack, output.FinalStack);
		assert.Equal(truthy, output.Truthy);
	}
}