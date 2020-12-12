package insertionsort

// insertionsort sort a slice of ints by traversing and swapping the elements
func insertionsort(target []int) {
	for i := 1; i < len(target); i++ {

		//break the loop if the previous element is less than the current element or we reached the first element in the list
		for i > 0 {
			if target[i-1] < target[i] {
				break
			}

			//swap with the previous element in the list
			target[i-1], target[i] = target[i], target[i-1]
			i--
		}
	}
}
