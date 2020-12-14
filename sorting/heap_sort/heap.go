package heapsort

func heapsort(target []int) {
	n := len(target)
	sortIt(target, n)
}

func sortIt(target []int, n int) {

	// Build heap (rearrange slice)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(target, n, i)
	}

	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// move current root to end
		target[0], target[i] = target[i], target[0]
		// call max heapify on the reduced heap
		heapify(target, i, 0)
	}
}

func heapify(target []int, n int, i int) {

	// If the parent node is stored at index i,
	// the left child can be calculated by 2 * i + 1 and
	// right child by 2 * i + 2
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && target[left] > target[largest] {
		largest = left
	}

	if right < n && target[right] > target[largest] {
		largest = right
	}

	if largest != i {
		target[i], target[largest] = target[largest], target[i]
		// recursively heapify the affected sub-tree
		heapify(target, n, largest)
	}
}
