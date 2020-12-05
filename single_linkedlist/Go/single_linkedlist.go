/*
Package singlelinkedlist implements:
	Create a a new Linkedlist and Node
	Insert a Node at the beginning of the list
	Insert a Node at the end of the list
	Delete a Node from the beginning of the list
	Delete a Node from the end of the list
	Print the Linkedlis
*/
package singlelinkedlist

type node struct {
	val  int
	Next *node
}

type linkedlist struct {
	length int
	Head   *node
}

//CreateList defines new linked list
func CreateList() *linkedlist {
	return &linkedlist{}
}

func newnode(val int) *node {
	return &node{val, nil}
}

//AddAtBeg adds a Node at the beginning of the list
func (ll *linkedlist) AddAtBeg(val int) {
	n := newnode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++

}

//AddAtEnd adds a Node at the end of the list
func (ll *linkedlist) AddAtEnd(val int) {
	n := newnode(val)

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
func (ll *linkedlist) DelAtBeg() int {
	if ll.Head == nil {
		return -1
	}
	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.val
}

//DelAtEnd deletes a node from the end of the list
func (ll *linkedlist) DelAtEnd() int {
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

func (ll *linkedlist) DeleteWithValute(val int) int {
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
