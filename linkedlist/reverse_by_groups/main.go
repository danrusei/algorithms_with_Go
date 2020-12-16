package main

import (
	"fmt"

	"github.com/danrusei/algorithms_with_go/linkedlist"
)

type llist struct {
	linkedlist.Linkedlist
}

// Reverses the linked list in groups of size k and returns the pointer to the new head node.
func (ll *llist) reverseGroups(node *linkedlist.Node, k int) *linkedlist.Node {
	curr := node
	next := &linkedlist.Node{}
	prev := &linkedlist.Node{}
	count := 0

	// reverse first k nodes of the linked list
	for {
		if curr == nil || count > k {
			break
		}
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}

	// next is now a pointer to (k+1)th node
	// Recursively call for the list starting from current.
	// And make rest of the list as next of first node
	if next != nil {
		ll.Head.Next = ll.reverseGroups(next, k)
	}

	// prev is new head of the input list
	return prev
}

func main() {
	list := llist{}
	list.AddAtEnd(9)
	list.AddAtEnd(8)
	list.AddAtEnd(7)
	list.AddAtEnd(6)
	list.AddAtEnd(5)
	list.AddAtEnd(4)
	list.AddAtEnd(3)
	list.AddAtEnd(2)
	list.AddAtEnd(1)
	fmt.Println(list.String())

	list.Head = list.reverseGroups(list.Head, 3)
	fmt.Println(list.String())

}
