package script

type Frame []byte;
type Stack []Frame;

func NewStack ()*Stack {
	return &Stack{};
}

func (s *Stack) Len() int {
	return len((*s));
}

func (s *Stack) Pop() Frame {
	x := (*s)[s.Len()-1];
	temp := (*s)[:s.Len()-1];
	s = &temp;

	return x;
}

func (s *Stack) Push(a Frame) {
	temp := append((*s), a);
	s = &temp;
}

func (s *Stack) Splice(n int, deleteCount int, b []Frame) {
	
}

func (s *Frame) Copy() Frame {
	temp := make([]byte, len((*s)));
	copy(temp,*s);

	return temp;
}

func (f *Frame) Int() int {
	return 0;
}

func (f *Frame) IsZero() bool {
	// Check if each byte in frame is 0.
	for _, byteInd := range *f {
		if byteInd != 0 {
			return false;
		}
	}

	return true;
}