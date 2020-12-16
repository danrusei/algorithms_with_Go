package main

import (
	"fmt"

	"github.com/danrusei/algorithms_with_go/linkedlist"
)

type llist struct {
	linkedlist.Linkedlist
}

// method to insert nodes, ordered by the values
func (ll *llist) sortedInsert(n *linkedlist.Node) {

	//special case for the head end
	if ll.Head == nil || ll.Head.Val >= n.Val {
		n.Next = ll.Head
		ll.Head = n
		ll.Length++
		return
	}

	//locate the node before the point of insertion
	cur := ll.Head
	for ; cur.Next != nil && cur.Next.Val < n.Val; cur = cur.Next {
	}
	n.Next = cur.Next
	cur.Next = n
	ll.Length++
}

func main() {
	list := llist{}
	node := linkedlist.NewNode(5)
	list.sortedInsert(node)
	node = linkedlist.NewNode(8)
	list.sortedInsert(node)
	node = linkedlist.NewNode(10)
	list.sortedInsert(node)
	node = linkedlist.NewNode(3)
	list.sortedInsert(node)
	node = linkedlist.NewNode(1)
	list.sortedInsert(node)
	node = linkedlist.NewNode(9)
	list.sortedInsert(node)

	fmt.Println(list.String())
}
