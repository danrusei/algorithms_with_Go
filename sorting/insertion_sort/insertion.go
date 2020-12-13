package insertionsort

// sorts the slice by traversing and swapping the elements
func insertionsort(target []int) {
	for i := 1; i < len(target); i++ {

		// break the loop if the previous element is less than the current element
		// or if it reaches the first element of the slice
		for i > 0 {
			if target[i-1] < target[i] {
				break
			}

			// swap the elements in the slice
			target[i-1], target[i] = target[i], target[i-1]
			i--
		}
	}
}
