package compare

import (
	"testing"
)

// this function creates the lists,
// the way the linkedlist is defined it allows us to check the length
func createLinkedList(vals []string) *Linkedlist {
	list := NewLinkedlist()
	node := NewNode(vals[0])
	list.Head = node
	list.Length = 1
	cur := list.Head
	for i, val := range vals {
		if i == 0 {
			continue
		}
		node := NewNode(val)
		cur.Next = node
		list.Length++
	}
	return list
}

func TestComparedStrings(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			list1 := createLinkedList(tc.list1)
			list2 := createLinkedList(tc.list2)
			result := compareStrings(list1, list2)
			if result != tc.result {
				t.Errorf("got: %v, want: %v", result, tc.result)
			}
		})
	}
}
