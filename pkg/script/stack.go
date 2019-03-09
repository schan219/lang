package script

type Frame []byte
type Stack []Frame

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Push(v int) Stack {
    return append(s, v)
}

func (s Stack) Pop() (Stack, int) {
	if (s.Len() == 0) {
		panic("Trying to pop from stack!")
	}

    l := len(s)
    return  s[:l-1], s[l-1]
}

func (s Frame) Copy() Frame {
	temp := make([]byte, len((*s)))
	copy(temp, s)

	return Frame(temp)
}

func (f Frame) Int() int {
	if len(f) > 4 {
		panic("Integer overflow!")
	}

	// Construct the int in little endian form.
	total := 0
	for i := 0; i < len(f); i++ {
		total |= (f[i] << (i * 8))
	}

	

	return total
}

func (f Frame) IsZero() bool {
	// Check if each byte in frame is 0.
	for _, byteInd := range *f {
		if byteInd != 0 {
			return false
		}
	}

	return true;
}