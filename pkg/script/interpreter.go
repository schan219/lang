package script

import (
	"errors"
)

var SUPPORTED_COINS = map[string]bool{
	"BSV": true,
	"BTC": true,
	"BCH": true,
	"ALL": true,
}

type Interpreter struct {
	CurrentScript []byte
	Stack	      [][]byte
	Coin          string
} 

func NewInterpreter(coinType string) (error, *Interpreter) {
	// Assert it's a valid coin.
	if _,ok := SUPPORTED_COINS[coinType]; !ok {
		return errors.New("Coin not supported, use BTC, BCH, BSV, or ALL"), nil
	}

	return nil, &Interpreter{coin: coinType}
}

func (intp *Interpreter) ExecInOut (input []byte, output []byte) (error, bool) {
	return intp.Exec(append(input, output...))
}

func (intp *Interpreter) ExecStack (stack [][]byte, script []byte) (error, bool) {
	intp.Stack = stack
	return intp.Exec(script);
}

func (intp *Interpreter) Exec(script []byte) (error, bool) {

	for {
		// Pop out the first command
		command, script := script[0], script[1:]
		
		// Attempt to execute the command!
		OP_FUNCS[command](intp,&script, command, &intp.Stack)
	}
}