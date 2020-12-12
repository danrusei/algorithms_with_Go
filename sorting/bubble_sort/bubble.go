package bubblesort

func bubblesort(target []int) {
	for firstnumber := 0; firstnumber < len(target); firstnumber++ {
		swapped := false
		for i := 0; i < len(target)-1-firstnumber; i++ {
			if target[i] > target[i+1] {
				target[i], target[i+1] = target[i+1], target[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
