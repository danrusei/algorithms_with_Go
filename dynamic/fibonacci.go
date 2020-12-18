package main

import "fmt"

// non-DP recursive fibonacci solution
func nonDPFibonacci(n int) int {
	if n < 2 {
		return n
	}
	return nonDPFibonacci(n-1) + nonDPFibonacci(n-2)
}

// DP Memoization fibonacci
func memFibonacci(n int) int {
	mem := make([]int, n+1)
	return calcFibonacci(mem, n)
}
func calcFibonacci(mem []int, n int) int {
	if n < 2 {
		return n
	}
	//if we have already solved this subproblem,
	//simply return the result from the cache
	if mem[n] != 0 {
		return mem[n]
	}

	mem[n] = calcFibonacci(mem, n-1) + calcFibonacci(mem, n-2)
	return mem[n]
}

//DP Tabulation fibonacci
func tabFibonacci(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Printf("nonDP recursive Fibonacci(10) = %d \n", nonDPFibonacci(10))
	fmt.Printf("DP Top-Down Fibonacci(10) = %d \n", memFibonacci(10))
	fmt.Printf("DP Bottom-Up Fibonacci(10) = %d \n", tabFibonacci(10))
}
