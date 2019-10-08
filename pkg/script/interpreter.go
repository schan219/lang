package interpreter

import (
	"github.com/bitcoinsv/bsvd/txscript"
)

// Executes a provided string ignoring CHECKSIG operations
func Execute(script []byte, starting [][]byte) {
	// TODO: Check for CHECKSIGS and complain.
	msgTx := NewMsgTx(1)
	msgTx.AddTxOut(&TxOut{
		Value: 0,
		PkScript: script
	})
	msgTx.LockTime = 0

	return txscript.NewEngine(script, msgTx, 0, nil, nil, 1).Execute()
}
