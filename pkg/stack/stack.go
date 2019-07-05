package parser

type Frame interface{}
type Stack []*Frame

func (s *Stack) Top() *Frame {
	return s[len(s) - 1]
}

func (s *Stack) Pop() *Frame {
	x := s.Top()
	*s = s[0:len(s) - 1]
	
	return x
}

func (s *Stack) Push(x Frame) {
	*s = append(s, x)
}
