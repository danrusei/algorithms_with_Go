/*
Package stack implements:
	a a new type named Stack, which is a slice of ints
	push method to add an element to the Stack
	pop method to remove the element from the Stac
*/
package stack

//Stack represent a stack
type Stack []int

//Push an elememnt on the top of the stack
func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

//Pop remove the top element from the stack
func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	lastindex := len(*s) - 1
	element := (*s)[lastindex]
	*s = (*s)[:lastindex]
	return element, true
}
