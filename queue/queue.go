/*
Package queue implements:
	a a new type named Stack, which is a slice of ints
	**Enqueue** method to add an element at the end of the Queue
	**Dequeue** method to remove the element from the the front of the Queue
*/
package queue

// Queue represent a stack
type Queue []int

// Enqueue insert an elememnt at the back of the queue
func (q *Queue) Enqueue(element int) {
	*q = append(*q, element)
}

// Dequeue remove an element at the front of the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}
