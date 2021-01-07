package main

import (
	"fmt"
	"sort"

	binarytree "github.com/danrusei/algorithms_with_go/binary_tree"
)

type comb struct {
	value int
	level int
}

//recursively traverse the binary tree and append in a map the element with the column as the key and a struct as the value
//the value contains the node value and its level in horizontal hierarchy
func findbottom(root *binarytree.Tree, m map[int]([]comb), col int, level int) {
	if root == nil {
		return
	}
	c := comb{root.Value, level}
	m[col] = append(m[col], c)
	findbottom(root.Left, m, col-1, level+1)
	findbottom(root.Right, m, col+1, level+1)
}

func bottomView(root *binarytree.Tree) {
	m := map[int]([]comb){}
	findbottom(root, m, 0, 0)

	//store the keys in a slice in order to access map elements in sorted order
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	//access the sorted map elements and print the bottom element from each column
	for _, k := range keys {
		var value int
		maxLevel := 0
		for _, c := range m[k] {
			if c.level > maxLevel {
				maxLevel = c.level
				value = c.value
			}

		}
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func main() {

	tree := binarytree.NewTree(20)
	tree.Left = binarytree.NewTree(8)
	tree.Right = binarytree.NewTree(22)
	tree.Left.Left = binarytree.NewTree(5)
	tree.Left.Right = binarytree.NewTree(3)
	tree.Right.Right = binarytree.NewTree(25)
	tree.Left.Right.Left = binarytree.NewTree(10)
	tree.Left.Right.Right = binarytree.NewTree(14)

	bottomView(tree)
}
