package subsetsum

func subsetsum(input []int, sum int) bool {

	n := len(input)

	// T[i][j] stores true if subset with sum j can be attained with
	// using items up to first i items
	mat := make([][]bool, n+1)
	for i := range mat {
		mat[i] = make([]bool, sum+1)
	}

	// if 0 items in the list and sum is non-zero
	for j := 1; j <= sum; j++ {
		mat[0][j] = false
	}

	// if sum is zero
	for i := 0; i <= n; i++ {
		mat[i][0] = true
	}

	// do for ith item
	for i := 1; i <= n; i++ {
		// consider all sum from 1 to sum
		for j := 1; j <= sum; j++ {
			// don't include ith element if j-input[i-1] is negative
			if input[i-1] > j {
				mat[i][j] = mat[i-1][j]
			} else {
				// find subset with sum j by excluding or including the ith item
				mat[i][j] = mat[i-1][j] || mat[i-1][j-input[i-1]]
			}
		}
	}

	// return maximum value
	return mat[n][sum]

}
