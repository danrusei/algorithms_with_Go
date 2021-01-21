package zigzag

func zigzag(input []int) {

	// if `less` is true indicate "<", else ">"
	// initially is true as we start with "<"
	// a < b > c < d > e < f
	less := true

	for i := 0; i <= len(input)-2; i++ {
		if less {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		} else {
			if input[i] < input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}

		// change flag to false, as we should operate on oposite signo
		less = !less
	}
}
