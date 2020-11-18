package queue

var testcases = []struct {
	description string
	queue       Queue
	element     int
	enqueue     Queue
	dequeue     int
	ok          bool
}{
	{
		description: "4 elements in queue",
		queue:       []int{1, 2, 3, 4},
		element:     5,
		enqueue:     []int{1, 2, 3, 4, 5},
		dequeue:     1,
		ok:          true,
	},
	{
		description: "empty queue",
		queue:       []int{},
		element:     5,
		enqueue:     []int{5},
		dequeue:     0,
		ok:          false,
	},
}
