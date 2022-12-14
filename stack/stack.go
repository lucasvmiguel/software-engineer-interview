package stack

type Stack struct {
	Values []string
	Size   int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value string) {
	s.Values = append(s.Values, value)
	s.Size++
}

func (s *Stack) Pop() string {
	if s.Size <= 0 {
		return ""
	}

	value := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	s.Size--
	return value
}
