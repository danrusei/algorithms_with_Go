package smallest

func smallestSlice(num int, input []int) int {

	//Initilize length of smallest slice as len(input) +1, which is not valid
	min := len(input) + 1

	for i := 0; i < len(input)-1; i++ {
		//initialize the sum with first element
		sum := input[i]

		//if first element itself is greater
		if sum > num {
			return 1
		}

		//iterate and add to the sum until the sum is bigger than the input number
		for j := i + 1; j < len(input)-1; j++ {
			sum += input[j]

			if sum > num {
				if j-i+1 < min {
					// save the new min
					min = j - i + 1
				}
				break
			}

		}
	}
	return min
}
