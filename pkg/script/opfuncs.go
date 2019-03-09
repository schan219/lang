package script

import (
	"encoding/binary"
	"errors"
)

var OP_FUNCS = map[int]interface{} {
	OP_0: pushDataLen(0),
	OP_PUSHDATA1: pushDataWidth(1),
	OP_PUSHDATA2: pushDataWidth(2),
	OP_PUSHDATA4: pushDataWidth(4),
}


func pushDataWidth (width int) interface{} {
	return func (stack Stack, cmd int, script []byte) (Stack, []byte, error) {
		var total int64; 
		var frameLen []byte;

		// We want to push the next n values.
		// The script must have that many bytes left.
		if len(script) < width {
			return nil, nil, errors.New("Not enough bytes left for pushdata size..");
		}


		// This the number of bytes we need to push in.
		var bytesLen = (cmd - OP_PUSHDATA1) + 1;
		// Little endian ftw!
		for i := 0; i < bytesLen; i++ {
			total += int64(script[0]) << uint(i*8);
			script = script[1:];
		}

		// Validate.
		if total > int64(len(script)) {
			panic("Not enough data on script for pushdata");
		}

		// Slice and push
		stack = stack.Push(Frame(script[0:total]))
		script = script[total:]

		return stack, script, nil;
	}
}

func pushDataLen(length int) interface{} {
	return func (stack *Stack, cmd int, script []byte) (*Stack, []byte, error) {
		if length > len(script) {
			panic("Not enough bytes ")
		}

		var frame []byte;
		frame, script = script[0:len(script)], script[len(script):];
		stack = stack.Push(frame);
		return stack, script, nil;
	}
}

func pushDataValue(value int) interface{} {
	return func (stack *Stack, cmd int, script []byte) (*Stack, []byte, error) {
		stack = stack.Push([]byte{byte(value)});
		return stack, script, nil;
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
		OP_FUNCS[opcode] = pushDataLen(opcode);
	}

	// We go from OP_1, to OP_16 inclusize
	for value := 1; value <= 16; value++ {
		// Since we start at 1, OP_1+1-1 == OP_1.
		OP_FUNCS[OP_1 + value - 1] = pushDataValue(value); 
	}

	OP_FUNCS[OP_DUP] = func (stack *Stack, cmd int, script []byte) (*Stack, []byte, error) {
		// b -- b b
		var temp1 Frame;
		temp1,stack = stack.Pop();
		temp2 := temp1.Copy();

		stack = stack.Push(temp1);
		stack = stack.Push(temp2);
		return stack, script, nil;
	}

	OP_FUNCS[OP_2DUP] = func (stack *Stack, cmd int, script []byte) (*Stack, []byte, error) {
		// a b -- a b a b
		b,stack := stack.Pop();
		a,stack := stack.Pop();

		stack = stack.Push(a.Copy());
		stack = stack.Push(b.Copy());
		stack = stack.Push(a.Copy());
		stack = stack.Push(b.Copy());
		return stack, script, nil;
	}

	OP_FUNCS[OP_3DUP] = func (stack *Stack, cmd int, script []byte) (*Stack, []byte, error) {
		// a b c -- a b c a b c
		b,stack := stack.Pop();
		a,stack := stack.Pop();
		c,stack := stack.Pop();

		stack = stack.Push(a.Copy());
		stack = stack.Push(b.Copy());
		stack = stack.Push(c.Copy());
		stack = stack.Push(a.Copy());
		stack = stack.Push(b.Copy());
		stack = stack.Push(c.Copy());
		return stack, script, nil;
	}

	OP_FUNCS[OP_CAT] = func (stack *Stack, cmd int, script []byte) (*Stack, []byte, error) {
		// Not inlined to prevent ambiguity.
		temp1,stack := stack.Pop();
		temp2,stack := stack.Pop();
		temp3 := append(temp2, temp1...);

		stack = stack.Push(temp3);
		return stack, script, nil;
	}

	OP_FUNCS[OP_PICK] = func (stack *Stack, cmd int, script []byte) (*Stack, []byte, error) {
		//Pop out the last number.
		topFrame, stack := stack.Pop();
		n := topFrame.Int();

		if (n < 0 || n > stack.Len()) {
			return nil, nil, errors.New("OP_PICK is trying to pick out of range!")
		}

		// pop to the top. conversion will throw.
		value := topFrame.Int();
		// Data we picked out.
		data := (*stack)[stack.Len() - value];

		if cmd == OP_ROLL {
			// If we ROLL, then we remove too.
			(*stack).Splice(value, -1, []Frame{})
		}

		// Push our data to the top of the stack
		stack = stack.Push(data);
		return stack, script, nil;
	}
	// OP_PICK / OP_ROLL are super similar.
	OP_FUNCS[OP_ROLL] = OP_FUNCS[OP_PICK];
}