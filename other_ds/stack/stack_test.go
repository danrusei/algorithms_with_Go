package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	for _, tc := range testcases {
		tc.stack.Push(tc.element)

		if len(tc.stack) != len(tc.push) {
			t.Fatalf("Fail, the lengths between current stack and expected are different")
		}

		for i := range tc.stack {
			if tc.stack[i] != tc.push[i] {
				t.Fatalf("Fail, expected: %v, got: %v", tc.push, tc.stack)
			}
		}
	}
}

func TestPop(t *testing.T) {
	for _, tc := range testcases {
		element, ok := tc.stack.Pop()
		if ok != tc.ok {
			t.Fatalf("Fail, got: %v, expected: %v", ok, tc.ok)
		}

		if element != tc.pop {
			t.Fatalf("Fail got: %d, expected: %d", element, tc.pop)
		}
	}
}
