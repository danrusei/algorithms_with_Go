package main

type node struct {
	val  int
	Next *node
}

type linkedlist struct {
	length int
	Head   *node
}

//CreateList defines new linked list
func NewLinkedlist() *linkedlist {
	return &linkedlist{}
}

func NewNode(val int) *node {
	return &node{val, nil}
}

func (llist *linkedlist) sortedInsert(node *node) {

}

func main() {
	llist := NewLinkedlist()
	newNode := NewNode(5)
	llist.sortedInsert(newNode)

}
