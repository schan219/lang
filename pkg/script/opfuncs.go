package script

import (
	"errors"
	"math/big"
)

var OP_FUNCS = map[int]interface{} {
	OP_0:         pushData(-1,1),

	OP_PUSHDATA1: pushData(1,-1),
	OP_PUSHDATA2: pushData(2,-1),
	OP_PUSHDATA4: pushData(4,-1),
}


func pushData (width int, value int) interface{} {
	// We just return the next n 
	if (value == -1) {
		return func () {
			
		}
	} else {
		return func () {

		}
	}
}

func init () {
	// we go from OP_DATA_1, to OP_DATA_75
	for opcode := OP_DATA_1; opcode <= OP_DATA_75; opcode++ {
		OP_FUNCS[opcode] = pushData(opcode, -1);
	}

	// We go from OP_1, to OP_16 inclusize
	for value := 1; value <= 16; value++ {
		// Since we start at 1, OP_1+1-1 == OP_1.
		OP_FUNCS[OP_1 + value - 1] = pushData(-1, value); 
	}



	OP_FUNCS[OP_PICK] = func (stack *Stack, command int) {
		//Pop out the last number.
		topFrame :=stack.Pop();
		n := topFrame.Int();

		if (n < 0 || n > stack.Len()) {
			panic("OP_PICK is trying to pick out of range!");
		}

		if command == OP_ROLL {

		}
	}
	// OP_PICK / OP_ROLL are super similar.
	OP_FUNCS[OP_ROLL] = OP_FUNCS[OP_PICK];
}