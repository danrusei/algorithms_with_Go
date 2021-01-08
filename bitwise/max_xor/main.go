package main

import (
	"fmt"
	"math"
)

func maxSubarrayXOR(numbers []int) int {
	// Initialize result
	response := math.MinInt32

	for i := range numbers {

		//store xor of current subarray
		currXor := 0

		// Pick ending points of subarrays starting with i
		for j := i; j < len(numbers); j++ {
			currXor = currXor ^ numbers[j]
			if response < currXor {
				response = currXor
			}
		}
	}
	return response
}

func main() {
	numbers := []int{8, 1, 2, 12, 7, 6}
	result := maxSubarrayXOR(numbers)
	fmt.Printf("Max subarray XOR is %d\n", result)
}
