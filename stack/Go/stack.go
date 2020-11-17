package stack

//Stack represent a stack
type Stack struct {
	items []int
}

//Push an elememnt on the top of the stack
func (s *Stack) Push(element int) {
	s.items = append(s.items, element)
}

//Pop remove the top element from the stack
func (s *Stack) Pop() int {
	lastindex := len(s.items) - 1
	element := s.items[lastindex]
	s.items = s.items[:lastindex]
	return element
}
