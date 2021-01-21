package main

import (
	"fmt"

	binarytree "github.com/danrusei/algorithms_with_go/binary_tree"
)

func removeShortPath(tree *binarytree.Tree, level int, k int) *binarytree.Tree {
	if tree == nil {
		return nil
	}

	// Traverse the tree in postorder fashion so that if a leaf
	// node path length is shorter than k, then that node and
	// all of its descendants till the node which are not
	// on some other path are removed.
	tree.Left = removeShortPath(tree.Left, level+1, k)
	tree.Right = removeShortPath(tree.Right, level+1, k)

	if tree.Left == nil && tree.Right == nil && level < k {
		tree = nil
	}

	return tree
}

func main() {

	tree := binarytree.NewTree(1)
	tree.Left = binarytree.NewTree(2)
	tree.Right = binarytree.NewTree(3)
	tree.Left.Left = binarytree.NewTree(4)
	tree.Left.Right = binarytree.NewTree(5)
	tree.Left.Left.Left = binarytree.NewTree(7)
	tree.Right.Right = binarytree.NewTree(6)
	tree.Right.Right.Left = binarytree.NewTree(8)

	fmt.Println("In order traversal of the original tree:")
	fmt.Println(tree)

	// this is the distance of the shortest path
	k := 4

	modifiedTree := removeShortPath(tree, 1, k)

	fmt.Println("In order traversal of the modified tree:")
	fmt.Println(modifiedTree)

}
