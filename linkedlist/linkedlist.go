/*
Package linkedlist implements:
	Create a a new Linkedlist and Node
	Insert a Node at the beginning of the list
	Insert a Node at the end of the list
	Delete a Node from the beginning of the list
	Delete a Node from the end of the list
	Print the Linkedlis
*/
package linkedlist

import "fmt"

// Node struct holds a node
type Node struct {
	val  int
	Next *Node
}

// Linkedlist struct holds a linked list
type Linkedlist struct {
	length int
	Head   *Node
}

// NewLinkedlist creates a new linked list
func NewLinkedlist() *Linkedlist {
	return &Linkedlist{}
}

// NewNode creates a new node
func NewNode(val int) *Node {
	return &Node{val, nil}
}

// utility function to print the linkedlist
func (ll *Linkedlist) String() string {
	cur := ll.Head
	printedlist := fmt.Sprintf("%d", cur.val)
	for ; cur.Next != nil; cur = cur.Next {
		printedlist = printedlist + fmt.Sprintf("%d \t", cur.val)
	}
	return printedlist

}

//AddAtBeg adds a Node at the beginning of the list
func (ll *Linkedlist) AddAtBeg(val int) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++

}

//AddAtEnd adds a Node at the end of the list
func (ll *Linkedlist) AddAtEnd(val int) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.length++
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}

	cur.Next = n
	ll.length++
}

//DelAtBeg deletes a node from the beginning of the list
func (ll *Linkedlist) DelAtBeg() int {
	if ll.Head == nil {
		return -1
	}
	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.val
}

//DelAtEnd deletes a node from the end of the list
func (ll *Linkedlist) DelAtEnd() int {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	val := cur.Next.val
	cur.Next = nil
	ll.length--

	return val
}

//DeleteWithValute deletes a node which value is equal to the function parameter
func (ll *Linkedlist) DeleteWithValute(val int) int {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	for cur.Next.val != val {
		if cur.Next.Next == nil {
			return -1
		}
		cur = cur.Next
	}
	value := cur.Next.val
	cur.Next = cur.Next.Next
	ll.length--

	return value
}
