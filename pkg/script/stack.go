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

func (f *Frame) Int() int {
	return 0;
}