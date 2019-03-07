package script

var OP_FUNCS = map[int]interface{} {
	OP_0:         pushData(-1,1),

	OP_PUSHDATA1: pushData(1,-1),
	OP_PUSHDATA2: pushData(2,-1),
	OP_PUSHDATA4: pushData(4,-1),
}


func pushData (width int, value int) interface{} {
	// We just return the next n 
	if (value == -1) {
		return func (stack *Stack, command int, script []byte) {
			
		}
	} else {
		return func (stack *Stack, command int, script []byte) {

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

	OP_FUNCS[OP_DUP] = func (stack *Stack, command int, script []byte) {
		// b -- b b
		temp1 := stack.Pop();
		temp2 := temp1.Copy();

		stack.Push(temp1);
		stack.Push(temp2);		
	}

	OP_FUNCS[OP_2DUP] = func (stack *Stack, command int, script []byte) {
		// a b -- a b a b
		b := stack.Pop();
		a := stack.Pop();

		stack.Push(a.Copy());
		stack.Push(b.Copy());
		stack.Push(a.Copy());
		stack.Push(b.Copy());
	}

	OP_FUNCS[OP_3DUP] = func (stack *Stack, command int, script []byte) {
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
	}

	OP_FUNCS[OP_CAT] = func (stack *Stack, command int, script []byte) {
		// Not inlined to prevent ambiguity.
		temp1 := stack.Pop();
		temp2 := stack.Pop();
		temp3 := append(temp2, temp1...);

		stack.Push(temp3);
	}

	OP_FUNCS[OP_PICK] = func (stack *Stack, command int, script []byte) {
		//Pop out the last number.
		topFrame :=stack.Pop();
		n := topFrame.Int();

		if (n < 0 || n > stack.Len()) {
			panic("OP_PICK is trying to pick out of range!");
		}

		// pop to the top. conversion will throw.
		value := topFrame.Int();
		// Data we picked out.
		data := (*stack)[stack.Len() - value];

		if command == OP_ROLL {
			// If we ROLL, then we remove too.
			stack.Splice(value, -1, []Frame{})
		}

		// Push our data to the top of the stack
		stack.Push(data);
	}
	// OP_PICK / OP_ROLL are super similar.
	OP_FUNCS[OP_ROLL] = OP_FUNCS[OP_PICK];
}