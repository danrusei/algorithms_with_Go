package compare

// Node struct holds a node
// We are not using this time the generic linked list defined at
// https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist
// as we need string values instead of int
type Node struct {
	Val  string
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
func NewNode(val string) *Node {
	return &Node{val, nil}
}
