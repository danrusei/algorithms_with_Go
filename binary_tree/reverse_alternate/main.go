package main

import (
	"fmt"

	binarytree "github.com/danrusei/algorithms_with_go/binary_tree"
)

// traverse tree and store odd level elements
func storeAlternate(tree *binarytree.Tree, store *[]int, level int) {
	if tree == nil {
		return
	}

	// Inorder traverse the tree and store only the values from alternate levels
	storeAlternate(tree.Left, store, level+1)

	if level%2 != 0 {
		*store = append(*store, tree.Value)
	}

	storeAlternate(tree.Right, store, level+1)
}

// traverse the tree and take the elements from the slice and store in odd level nodes
func modifyTree(tree *binarytree.Tree, store []int, index int, level int) int {

	if tree == nil {
		return index
	}

	// Inorder traverse the tree and update level values from the store if the level is odd
	index = modifyTree(tree.Left, store, index, level+1)

	if level%2 != 0 {
		tree.Value = store[index]
		index++
	}

	index = modifyTree(tree.Right, store, index, level+1)

	return index
}

func reverseAlternate(tree *binarytree.Tree) {

	//store all alternate levels nodes
	store := []int{}

	// inorder traverse and store the values from alternate levels
	storeAlternate(tree, &store, 0)

	// reverse the slice
	for i, j := 0, len(store)-1; i < j; i, j = i+1, j-1 {
		store[i], store[j] = store[j], store[i]
	}

	index := 0
	_ = modifyTree(tree, store, index, 0)
}

func main() {

	tree := binarytree.NewTree(11)
	tree.Left = binarytree.NewTree(12)
	tree.Right = binarytree.NewTree(13)
	tree.Left.Left = binarytree.NewTree(14)
	tree.Left.Left.Left = binarytree.NewTree(1)
	tree.Left.Left.Right = binarytree.NewTree(2)
	tree.Left.Right = binarytree.NewTree(15)
	tree.Left.Right.Left = binarytree.NewTree(3)
	tree.Left.Right.Right = binarytree.NewTree(4)
	tree.Right.Left = binarytree.NewTree(16)
	tree.Right.Left.Left = binarytree.NewTree(5)
	tree.Right.Left.Right = binarytree.NewTree(6)
	tree.Right.Right = binarytree.NewTree(16)
	tree.Right.Right.Left = binarytree.NewTree(7)
	tree.Right.Right.Right = binarytree.NewTree(8)

	fmt.Println("In order traversal of the original tree:")
	fmt.Println(tree)

	reverseAlternate(tree)

	fmt.Println("In order traversal of the modified tree:")
	fmt.Println(tree)

}
