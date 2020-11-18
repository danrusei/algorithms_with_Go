package queue

//Queue represent a stack
type Queue []int

//Enqueue insert an elememnt at the back of the queue
func (q *Queue) Enqueue(element int) {
	*q = append(*q, element)
}

//Dequeue remove an element at the front of the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}
