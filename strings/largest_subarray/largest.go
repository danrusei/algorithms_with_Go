package largest

func largestSlice(input []int) int {

	maxLen := 1
	for i := 0; i < len(input)-1; i++ {
		//  Initialize min and max for all slices starting with i
		minN := input[i]
		maxN := input[i]

		// Consider all subarrays starting with i and ending with j
		for j := i + 1; j < len(input); j++ {
			if input[j] < minN {
				minN = input[j]
			}
			if input[j] > maxN {
				maxN = input[j]
			}

			// calculate if the current slice has all contiguous elements
			if maxN-minN == j-i {
				if maxLen < maxN-minN+1 {
					maxLen = maxN - minN + 1
				}
			}
		}
	}
	return maxLen
}
