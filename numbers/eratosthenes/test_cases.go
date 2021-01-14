package primes

var testcases = []struct {
	name     string
	num      int
	expected []int
}{
	{
		name:     "10_value",
		num:      10,
		expected: []int{2, 3, 5, 7},
	},
	{
		name:     "20_value",
		num:      20,
		expected: []int{2, 3, 5, 7, 11, 13, 17, 19},
	},
	{
		name:     "99_value",
		num:      99,
		expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
	},
}
