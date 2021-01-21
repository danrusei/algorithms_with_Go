package zigzag

var testcases = []struct {
	name   string
	input  []int
	output []int
}{
	{
		name:   "first_10",
		input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		output: []int{1, 3, 2, 5, 4, 7, 6, 9, 8, 10},
	},
	{
		name:   "7_values",
		input:  []int{4, 3, 7, 8, 6, 2, 1},
		output: []int{3, 7, 4, 8, 2, 6, 1},
	},
	{
		name:   "4_values",
		input:  []int{1, 4, 3, 2},
		output: []int{1, 4, 2, 3},
	},
}
