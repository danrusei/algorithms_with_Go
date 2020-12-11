package bubblesort

var testcases = []struct {
	description string
	original    []int
	sorted      []int
}{
	{
		description: "4 elements in list",
		original:    []int{4, 2, 3, 1},
		sorted:      []int{1, 2, 3, 4},
	},
	{
		description: "empty list",
		original:    []int{},
		sorted:      []int{},
	},
	{
		description: "already sorted list",
		original:    []int{10, 20, 30, 40},
		sorted:      []int{10, 20, 30, 40},
	},
	{
		description: "9 elements in the list",
		original:    []int{11, 21, 12, 4, 3, 100, 1000, 1, 123},
		sorted:      []int{1, 3, 4, 11, 12, 21, 100, 123, 1000},
	},
}
