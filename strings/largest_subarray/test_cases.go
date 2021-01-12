package largest

var testcases = []struct {
	name   string
	input  []int
	output int
}{
	{
		name:   "entire_slice",
		input:  []int{10, 12, 11},
		output: 3,
	},
	{
		name:   "middle",
		input:  []int{14, 12, 11, 20},
		output: 2,
	},
	{
		name:   "larger_slice",
		input:  []int{1, 56, 58, 57, 90, 92, 94, 93, 91, 45},
		output: 5,
	},
}
