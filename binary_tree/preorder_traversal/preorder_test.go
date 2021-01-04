package preorder

import (
	"fmt"
	"reflect"
	"testing"

	binarytree "github.com/danrusei/algorithms_with_go/binary_tree"
)

func createBinaryTree(treeVals []int) *binarytree.Tree {

	tree := binarytree.NewTree(treeVals[0])

	for i, num := range treeVals {
		if i == 0 {
			continue
		}
		tree.Insert(num)
	}
	return tree
}

func TestPreorderTraversal(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tree := createBinaryTree(tc.tree)
			traverse := []int{}
			tree.PreOrderTraverse(func(elem int) {
				traverse = append(traverse, elem)
			})
			fmt.Printf("Input: %v\n", tc.traversal)
			fmt.Printf("Preorder traversal: %v\n", traverse)
			if result := reflect.DeepEqual(tc.traversal, traverse); result != tc.ok {
				t.Errorf("got: %v, want: %v", result, tc.ok)
			}
		})
	}
}
