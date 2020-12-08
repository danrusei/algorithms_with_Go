package main

import "fmt"

type node struct {
	val  int
	Next *node
}

type linkedlist struct {
	length int
	Head   *node
}

// utility function to print the linkedlist
func (ll *linkedlist) String() string {
	cur := ll.Head
	printedlist := fmt.Sprintf("%d", cur.val)
	for ; cur.Next != nil; cur = cur.Next {
		printedlist = printedlist + fmt.Sprintf("%d \t", cur.val)
	}
	return printedlist

}

// creates a new linkedlist
func newLinkedlist() *linkedlist {
	return &linkedlist{}
}

// creates a new node
func newNode(val int) *node {
	return &node{val, nil}
}

// method to insert nodes, ordered by the values
func (ll *linkedlist) sortedInsert(n *node) {

	//special case for the head end
	if ll.Head == nil || ll.Head.val >= n.val {
		n.Next = ll.Head
		ll.Head = n
		ll.length++
		return
	}

	//locate the node before the point of insertion
	cur := ll.Head
	for ; cur.Next != nil && cur.Next.val < n.val; cur = cur.Next {
	}
	n.Next = cur.Next
	cur.Next = n
	ll.length++
}

func main() {
	llist := newLinkedlist()
	node := newNode(5)
	llist.sortedInsert(node)
	node = newNode(8)
	llist.sortedInsert(node)
	node = newNode(10)
	llist.sortedInsert(node)
	node = newNode(3)
	llist.sortedInsert(node)
	node = newNode(1)
	llist.sortedInsert(node)
	node = newNode(9)
	llist.sortedInsert(node)

	fmt.Println(llist)
}
