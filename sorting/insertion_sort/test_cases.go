package insertionsort

var testcases = []struct {
	name     string
	original []int
	sorted   []int
}{
	{
		name:     "4 elements in list",
		original: []int{4, 2, 3, 1},
		sorted:   []int{1, 2, 3, 4},
	},
	{
		name:     "empty list",
		original: []int{},
		sorted:   []int{},
	},
	{
		name:     "already sorted list",
		original: []int{10, 20, 30, 40},
		sorted:   []int{10, 20, 30, 40},
	},
	{
		name:     "9 elements in the list",
		original: []int{11, 21, 12, 4, 3, 100, 1000, 1, 123},
		sorted:   []int{1, 3, 4, 11, 12, 21, 100, 123, 1000},
	},
}
