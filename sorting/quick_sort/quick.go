package quicksort

func quicksort(target []int) {
	sortIt(target, 0, len(target)-1)
}

//recursively sort the sub-slices
func sortIt(target []int, left int, right int) {

	// returns when there is nothing left to traverse
	if left >= right {
		return
	}

	//swap the elements and returns the dividing point between right and left side
	partIndex := partition(target, left, right)

	sortIt(target, left, partIndex-1)
	sortIt(target, partIndex, right)
}

func partition(target []int, left int, right int) int {

	//select the pivot point, could be randomly selected, in this case I take the middle element
	pivot := target[(left+right)/2]

	// loop over the slice untill it finds the elements to be swapped
	for left <= right {
		for target[left] < pivot {
			left++
		}
		for target[right] > pivot {
			right--
		}
		if left <= right {
			target[left], target[right] = target[right], target[left]
			left++
			right--
		}
	}
	return left
}
