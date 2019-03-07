package script

import (
	_"fmt"
	"testing"

	"lang/pkg/script"
	"github.com/stretchr/testify/assert"
)

type Frame script.Frame;
type Stack script.Stack;

type TestVal struct {
	FinalStack   *script.Stack
	IsError      bool
}

func TestMainPrimitives (t *testing.T) {
	assert := assert.New(t)
	err, intp := script.NewInterpreter("ALL");

	if err != nil {
		panic(err)
	}

	programs := map[string]TestVal {
		string([]byte{script.OP_0}): TestVal {
			FinalStack: script.NewStack(),
			IsError: false,
		},
		string([]byte{script.OP_1,script.OP_2,}): TestVal {
			FinalStack: script.NewStack(),
			IsError: false,
		},
		string([]byte{script.OP_5,script.OP_2,script.OP_3,script.OP_5,script.OP_7}): TestVal {
			FinalStack: script.NewStack(),
			IsError: false,
		},
	}

	for sourceCode, output := range programs {
		err, valid := intp.Exec([]byte(sourceCode));

		if output.IsError {
			assert.NotNil(err);
		} else {
			assert.Nil(err);
		}

		if output.FinalStack != nil {
			assert.True(valid);
		}
	}
}