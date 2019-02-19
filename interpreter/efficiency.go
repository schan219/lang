package interpreter

var efficiencyMap = map[[]int][]int{
	// Stack Operations
	[]int{OP_1, OP_PICK}:[]int{OP_DUP},
	[]int{OP_2, OP_PICK}:[]int{OP_OVER},
	[]int{OP_DROP,OP_DROP}:[]int{OP_2DROP},
	[]int{OP_2,OP_PICK,OP_2,OP_PICK}:[]int{OP_2DUP},
	[]int{OP_3,OP_PICK,OP_3,OP_PICK,OP_3,OP_PICK}:[]int{OP_3DUP},
	
	// Logical Operations
	[]int{OP_NOT,OP_IF}:[]int{OP_NOTIF},

	// Crypto operations
	[]int{OP_SHA256,OP_SHA256}:[]int{OP_HASH256},
	[]int{OP_SHA256,OP_RIPEMD160}:[]int{OP_HASH160},
	[]int{OP_CHECKSIG,OP_VERIFY}:[]int{OP_CHECKSIGVERIFY},
	[]int{OP_CHECKMULTISIG,OP_VERIFY}:[]int{OP_CHECKMULTISIGVERIFY},
}


func Optimize (code []byte) {
	
}