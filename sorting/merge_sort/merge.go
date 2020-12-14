package mergesort

func mergesort(target []int) {
	left := 0
	right := len(target) - 1
	sortIt(target, left, right)
}

// run recursively with a modified left and right index of the slice
func sortIt(target []int, left int, right int) {
	if left >= right {
		return
	}
	mid := (left + right - 1) / 2

	sortIt(target, left, mid)
	sortIt(target, mid+1, right)

	merge(target, left, mid, right)
}

func merge(target []int, left int, mid int, right int) {

	//find the sizes of the two slices
	n1 := mid - left + 1
	n2 := right - mid

	//create temp slices
	leftSlice := make([]int, n1)
	rightSlice := make([]int, n2)

	//copy data to the temp slices
	for i := 0; i < n1; i++ {
		leftSlice[i] = target[left+i]
	}
	for j := 0; j < n2; j++ {
		rightSlice[j] = target[mid+1+j]
	}

	//merge the temp slices
	var i, j = 0, 0
	k := left

	for i < n1 && j < n2 {
		if leftSlice[i] <= rightSlice[j] {
			target[k] = leftSlice[i]
			i++
		} else {
			target[k] = rightSlice[j]
			j++
		}
		k++
	}

	//copy the remaining elements from leftSlice
	for i < n1 {
		target[k] = leftSlice[i]
		i++
		k++
	}

	//copy the remaining elements from leftSlice
	for j < n2 {
		target[k] = rightSlice[j]
		j++
		k++
	}
}
