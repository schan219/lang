package script

import (
	"fmt"
	"errors"
)

var SUPPORTED_COINS = map[string]bool{
	//"BSV": true,
	//"BTC": true,
	//"BCH": true,
	"ALL": true,
}

type Interpreter struct {
	CurrentScript  []byte
	Stack	       *Stack
	Coin           string
	Invalid        bool
} 

func NewInterpreter(coinType string) (error, *Interpreter) {
	// Assert it's a valid coin.
	if _,ok := SUPPORTED_COINS[coinType]; !ok {
		return errors.New("Coin not supported, use BTC, BCH, BSV, or ALL"), nil
	}
	// Initia 
	InitFuncs();

	return nil, &Interpreter{Stack: NewStack(), Coin: coinType}
}

func (intp *Interpreter) Exec(script []byte) (error, bool) {
	// Put in some validation steps for different coin types
	for {
		// Finished reading the script!
		if len(script) == 0 {
			break;
		}
		var command byte;
		// Pop out the first command
		command, script = script[0], script[1:];
		
		// Attempt to execute the command!
		fn, exists := OP_FUNCS[int(command)]

		// Throw if weird opcode.
		if !exists {
			return errors.New(
				fmt.Sprintf("Unknown OpCode encountered: %x", command),
			), false;
		}

		// Execute the function..
		var err error;
		script, err = fn.(func(*Stack, int, []byte) ([]byte, error))(intp.Stack, int(command), script);

		if err != nil {
			return err, false;
		}
	}
	
	if intp.Invalid {
		return errors.New("Invalid transaction!"), false;
	}

	// If the top of the stack is a valid number return true, otherwise false.
	if intp.Stack.Len() > 0 {
		p := intp.Stack.Pop();
		return nil, p.IsZero();
	} else {
		return errors.New("Stack is empty"), false;
	}

}

func (intp *Interpreter) ExecInOut (input []byte, output []byte) (error, bool) {
	return intp.Exec(append(input, output...))
}

func (intp *Interpreter) ExecStack (stack *Stack, script []byte) (error, bool) {
	intp.Stack = stack;
	return intp.Exec(script);
}