package binarytree

import "fmt"

// Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

// NewTree instantiate the binary tree
func NewTree(k int) *Tree {
	return &Tree{
		Value: k,
		Left:  nil,
		Right: nil,
	}

}

// Insert a new value to the binary tree
func (t *Tree) Insert(val int) *Tree {
	if t == nil {
		return NewTree(val)
	}
	if val < t.Value {
		t.Left = t.Left.Insert(val)
	} else {
		t.Right = t.Right.Insert(val)
	}
	return t
}

// InOrderTraverse - the left subtree is visited first, then the root and later the right sub-tree.
func (t *Tree) InOrderTraverse(f func(int)) {
	if t != nil {
		t.Left.InOrderTraverse(f)
		f(t.Value)
		t.Right.InOrderTraverse(f)
	}
}

// PreOrderTraverse - the root node is visited first, then the left subtree and finally the right subtree.
func (t *Tree) PreOrderTraverse(f func(int)) {
	if t != nil {
		f(t.Value)
		t.Left.PreOrderTraverse(f)
		t.Right.PreOrderTraverse(f)
	}
}

// PostOrderTraverse - traverse the left subtree, then the right subtree and finally the root node
func (t *Tree) PostOrderTraverse(f func(int)) {
	if t != nil {
		t.Left.PostOrderTraverse(f)
		t.Right.PostOrderTraverse(f)
		f(t.Value)
	}
}

// Used to print the binary tree
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
