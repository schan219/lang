package script

import (
	"errors"
)

var OP_FUNCS = map[int]interface{} {
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