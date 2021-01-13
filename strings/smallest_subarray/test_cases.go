package smallest

var testcases = []struct {
	name   string
	num    int
	input  []int
	output int
}{
	{
		name:   "51_value",
		num:    51,
		input:  []int{1, 4, 45, 6, 0, 19},
		output: 3,
	},
	{
		name:   "9_value",
		num:    9,
		input:  []int{1, 10, 5, 2, 7},
		output: 1,
	},
	{
		name:   "280_value",
		num:    280,
		input:  []int{1, 11, 100, 1, 0, 200, 3, 2, 1, 250},
		output: 4,
	},
	{
		name:   "8_value",
		num:    8,
		input:  []int{1, 2, 4},
		output: 0,
	},
}
