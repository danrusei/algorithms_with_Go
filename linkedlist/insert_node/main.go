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

func (ll *linkedlist) String() string {
	cur := ll.Head
	printedlist := fmt.Sprintf("%d", cur.val)
	for ; cur.Next != nil; cur = cur.Next {
		printedlist = printedlist + fmt.Sprintf("%d \t", cur.val)
	}
	return printedlist

}

func newLinkedlist() *linkedlist {
	return &linkedlist{}
}

func newNode(val int) *node {
	return &node{val, nil}
}

func (ll *linkedlist) sortedInsert(n *node) {
	if ll.Head == nil || ll.Head.val >= n.val {
		n.Next = ll.Head
		ll.Head = n
		ll.length++
		return
	}

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
