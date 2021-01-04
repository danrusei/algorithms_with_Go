package preorder

var testcases = []struct {
	name      string
	tree      []int
	traversal []int
	ok        bool
}{
	{
		name:      "length_3_true",
		tree:      []int{2, 4, 3},
		traversal: []int{2, 4, 3},
		ok:        true,
	},
	{
		name:      "length_3_false",
		tree:      []int{2, 4, 1},
		traversal: []int{2, 4, 1},
		ok:        false,
	},
	{
		name:      "length_5_true",
		tree:      []int{40, 30, 35, 80, 100},
		traversal: []int{40, 30, 35, 80, 100},
		ok:        true,
	},
	{
		name:      "length_6_false",
		tree:      []int{40, 30, 35, 20, 80, 100},
		traversal: []int{40, 30, 35, 20, 80, 100},
		ok:        false,
	},
}
