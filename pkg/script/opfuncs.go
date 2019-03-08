package script

import (
	"encoding/binary"
	"errors"
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
		return func (stack *Stack, cmd int, script []byte) ([]byte, error) {
			// We want to push the next n values.
			// The script must have that many bytes left.
			if len(script) < width {
				return nil, errors.New("Not enough bytes left for pushdata size..");
			}

			var total int64; 
			var frameLen []byte;

			// Create the varint as a bytearray based on first byte
			switch width {
			case OP_PUSHDATA1:
				frameLen, script = script[0:2], script[2:];
				total = int64(binary.LittleEndian.Uint16(frameLen));
			case OP_PUSHDATA2:
				frameLen, script = script[0:4], script[4:];
				total = int64(binary.LittleEndian.Uint32(frameLen));
			case OP_PUSHDATA4:
				frameLen, script = script[0:8], script[8:];
				total = int64(binary.LittleEndian.Uint64(frameLen));
			default:
				return nil, errors.New("Invalid pushdata operation!");
			}

			// Validate.
			if total > int64(len(script)) {
				return nil, errors.New("Not enough data on stack for pushdata");
			}

			// Slice and push
			stack.Push(script[0:total])
			script = script[total:]

			return script, nil;
		}
	} else {
		return func (stack *Stack, cmd int, script []byte) ([]byte, error) {
			stack.Push([]byte{byte(value)});
			return script, nil;
		}
	}
}

// Singleton pattern ish, should be refactored in the future.
var fnInit = false;
func InitFuncs () {
	if fnInit {
		return;
	}
	fnInit = true;
	
	// we go from OP_DATA_1, to OP_DATA_75
	for opcode := OP_DATA_1; opcode <= OP_DATA_75; opcode++ {
		OP_FUNCS[opcode] = pushData(opcode, -1);
	}

	// We go from OP_1, to OP_16 inclusize
	for value := 1; value <= 16; value++ {
		// Since we start at 1, OP_1+1-1 == OP_1.
		OP_FUNCS[OP_1 + value - 1] = pushData(-1, value); 
	}

	OP_FUNCS[OP_DUP] = func (stack *Stack, cmd int, script []byte) ([]byte, error) {
		// b -- b b
		temp1 := stack.Pop();
		temp2 := temp1.Copy();

		stack.Push(temp1);
		stack.Push(temp2);
		return script, nil;
	}

	OP_FUNCS[OP_2DUP] = func (stack *Stack, cmd int, script []byte) ([]byte, error) {
		// a b -- a b a b
		b := stack.Pop();
		a := stack.Pop();

		stack.Push(a.Copy());
		stack.Push(b.Copy());
		stack.Push(a.Copy());
		stack.Push(b.Copy());
		return script, nil;
	}

	OP_FUNCS[OP_3DUP] = func (stack *Stack, cmd int, script []byte) ([]byte, error) {
		// a b c -- a b c a b c
		b := stack.Pop();
		a := stack.Pop();
		c := stack.Pop();

		stack.Push(a.Copy());
		stack.Push(b.Copy());
		stack.Push(c.Copy());
		stack.Push(a.Copy());
		stack.Push(b.Copy());
		stack.Push(c.Copy());
		return script, nil;
	}

	OP_FUNCS[OP_CAT] = func (stack *Stack, cmd int, script []byte) ([]byte, error) {
		// Not inlined to prevent ambiguity.
		temp1 := stack.Pop();
		temp2 := stack.Pop();
		temp3 := append(temp2, temp1...);

		stack.Push(temp3);
		return script, nil;
	}

	OP_FUNCS[OP_PICK] = func (stack *Stack, cmd int, script []byte) ([]byte, error) {
		//Pop out the last number.
		topFrame :=stack.Pop();
		n := topFrame.Int();

		if (n < 0 || n > stack.Len()) {
			return nil, errors.New("OP_PICK is trying to pick out of range!")
		}

		// pop to the top. conversion will throw.
		value := topFrame.Int();
		// Data we picked out.
		data := (*stack)[stack.Len() - value];

		if cmd == OP_ROLL {
			// If we ROLL, then we remove too.
			stack.Splice(value, -1, []Frame{})
		}

		// Push our data to the top of the stack
		stack.Push(data);
		return script, nil;
	}
	// OP_PICK / OP_ROLL are super similar.
	OP_FUNCS[OP_ROLL] = OP_FUNCS[OP_PICK];
}