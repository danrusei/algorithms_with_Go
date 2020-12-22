package subsetsum

var testcases = []struct {
	name   string
	input  []int
	sum    int
	output bool
}{
	{
		name:   "True as 4 + 5 = 9",
		input:  []int{3, 34, 4, 12, 5, 2},
		sum:    9,
		output: true,
	},
	{
		name:   "No sum of 2 elements from the list = 30",
		input:  []int{3, 34, 4, 12, 5, 2},
		sum:    30,
		output: false,
	},
}
