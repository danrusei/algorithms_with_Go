package binarytree

import "fmt"

func createTree() *Tree {
	bst := NewTree(5)
	bst.Insert(8).Insert(4).Insert(10).Insert(2).Insert(6).Insert(1).Insert(3).Insert(7).Insert(9)
	return bst
}

func ExampleTreeInsert() {

	bst := createTree()
	fmt.Println(bst)
	// Output:
	// ((((1) 2 (3)) 4) 5 ((6 (7)) 8 ((9) 10)))
}

func ExampleInOrderTraverse() {

	bst := createTree()
	bst.InOrderTraverse(func(i int) {
		fmt.Printf("%d\t", i)
	})
	// Output:
	// 1	2	3	4	5	6	7	8	9	10
}

func ExamplePreOrderTraverse() {

	bst := createTree()
	bst.PreOrderTraverse(func(i int) {
		fmt.Printf("%d\t", i)
	})
	// Output:
	// 5	4	2	1	3	8	6	7	10	9
}

func ExamplePostOrderTraverse() {

	bst := createTree()
	bst.PostOrderTraverse(func(i int) {
		fmt.Printf("%d\t", i)
	})
	// Output:
	// 1	3	2	4	7	6	9	10	8	5
}
