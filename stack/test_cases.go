package stack

var testcases = []struct {
	description string
	stack       Stack
	element     int
	push        Stack
	pop         int
	ok          bool
}{
	{
		description: "4 elements in stack",
		stack:       []int{1, 2, 3, 4},
		element:     5,
		push:        []int{1, 2, 3, 4, 5},
		pop:         4,
		ok:          true,
	},
	{
		description: "empty stack",
		stack:       []int{},
		element:     5,
		push:        []int{5},
		pop:         0,
		ok:          false,
	},
}
