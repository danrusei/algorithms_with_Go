package main

import (
	"fmt"
	"math"

	binarytree "github.com/danrusei/algorithms_with_go/binary_tree"
)

func minDepth(tree *binarytree.Tree) int {

	// Corner case. Should never be hit unless the code is called on nil tree
	if tree == nil {
		return 0
	}

	// Base case : Leaf Node. This accounts for height = 1.
	if tree.Left == nil && tree.Right == nil {
		return 1
	}

	// If left subtree is nil, recur for right subtree
	if tree.Left == nil {
		return minDepth(tree.Right) + 1
	}

	// If right subtree is nil, recur for left subtree
	if tree.Right == nil {
		return minDepth(tree.Left) + 1
	}

	return int(math.Min(float64(minDepth(tree.Left)), float64(minDepth(tree.Right))) + 1)

}

func main() {

	tree := binarytree.NewTree(1)
	tree.Left = binarytree.NewTree(2)
	tree.Right = binarytree.NewTree(3)
	tree.Left.Left = binarytree.NewTree(4)
	tree.Left.Right = binarytree.NewTree(5)

	result := minDepth(tree)

	fmt.Printf("The minimum depth of the binary tree is : %d\n", result)
}
