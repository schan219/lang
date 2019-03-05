package script

import (
	"errors"
)

var OP_FUNCS = map[int]interface{} {
	OP_0:         pushData(-1,1),

	OP_PUSHDATA1: pushData(1,-1),
	OP_PUSHDATA2: pushData(2,-1),
	OP_PUSHDATA4: pushData(4,-1),
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


func pushDataCmd (intp *Interpreter, script []byte, cmd byte) error {
	if int(cmd) > len(script) {
		return errors.New("Push being called with not enough elements")
	}

	// OP_FALSE exception
	if cmd == 0 {
		return intp.pushFrame([]byte{0})
	}

	frame, script := script[0:cmd], script[cmd:]
	return intp.pushFrame(frame)
}