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
// the value of the Node could be anything (int, string or a struncture)
// interface{} - an empty interface may hold values of any type.
// if using interface{} confuse you, change with any type you would like
// initially I started with int value, but some problems require strings as values 
// another example: implementing Hash-Table I need even to define a Key-Value struct
type Node struct {
	Val  interface{}
	Next *Node
}

// Linkedlist struct holds a linked list
type Linkedlist struct {
	Length int
	Head   *Node
}

// NewLinkedlist creates a new linked list
func NewLinkedlist() *Linkedlist {
	return &Linkedlist{}
}

// NewNode creates a new node
func NewNode(val interface{}) *Node {
	return &Node{val, nil}
}

// utility function to print the linkedlist
func (ll *Linkedlist) String() string {
	cur := ll.Head
	var printedlist string
	for ; cur.Next != nil; cur = cur.Next {
		printedlist = printedlist + fmt.Sprintf("%v \t", cur.Val)
	}
	printedlist = printedlist + fmt.Sprintf("%v", cur.Val)
	return printedlist
}

//AddAtBeg adds a Node at the beginning of the list
func (ll *Linkedlist) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.Length++

}

//AddAtEnd adds a Node at the end of the list
func (ll *Linkedlist) AddAtEnd(val interface{}) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.Length++
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}

	cur.Next = n
	ll.Length++
}

//DelAtBeg deletes a node from the beginning of the list
func (ll *Linkedlist) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}
	cur := ll.Head
	ll.Head = cur.Next
	ll.Length--

	return cur.Val
}

//DelAtEnd deletes a node from the end of the list
func (ll *Linkedlist) DelAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	val := cur.Next.Val
	cur.Next = nil
	ll.Length--

	return val
}

//DeleteWithValute deletes a node which value is equal to the function parameter
func (ll *Linkedlist) DeleteWithValute(val interface{}) interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	for cur.Next.Val != val {
		if cur.Next.Next == nil {
			return -1
		}
		cur = cur.Next
	}
	value := cur.Next.Val
	cur.Next = cur.Next.Next
	ll.Length--

	return value
}
