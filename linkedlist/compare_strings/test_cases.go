package compare

var testcases = []struct {
	name   string
	list1  []string
	list2  []string
	result int
}{
	{
		name:   "second_list_greater",
		list1:  []string{"g", "e", "e", "k", "s", "a"},
		list2:  []string{"g", "e", "e", "k", "s", "b"},
		result: -1,
	},
	{
		name:   "the_same_list",
		list1:  []string{"g", "e", "e", "k", "s"},
		list2:  []string{"g", "e", "e", "k", "s"},
		result: 0,
	},
	{
		name:   "first_list_greater",
		list1:  []string{"g", "e", "e", "k", "s", "c"},
		list2:  []string{"g", "e", "e", "k", "s"},
		result: 1,
	},
}
