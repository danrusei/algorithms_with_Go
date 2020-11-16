package stack

import (
	"testing"
)

var testcases = []struct {
	description string
	stack       Stack
	element     int
	push        Stack
	pop         int
}{
	{
		description: "some test",
		stack: Stack{
			items: []int{1, 2, 3, 4},
		},
		element: 5,
		push: Stack{
			items: []int{1, 2, 3, 4, 5},
		},
		pop: 4,
	},
	{
		description: "some test",
		stack: Stack{
			items: []int{1, 2, 3, 4},
		},
		element: 5,
		push: Stack{
			items: []int{1, 2, 3, 4, 5},
		},
		pop: 4,
	},
}

func TestPush(t *testing.T) {
	for _, tc := range testcases {
		tc.stack.Push(tc.element)
		if len(tc.stack.items) != len(tc.push.items) {
			t.Fatalf("Fail, expected: %v, got: %v", tc.push, tc.stack)
		}
	}
}
