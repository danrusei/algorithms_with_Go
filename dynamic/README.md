# Dynamic Programming

Dynamic Programming (DP) is an algorithmic technique for solving an optimization problem by breaking it down into simpler subproblems in a `recurisive` manner and utilizing the fact that the optimal solution to the overall problem depends upon the optimal solution to its subproblems.

It can be achieved in either of two ways:

* **Top-down approach (Memoization)**: This is the direct fall-out of the recursive formulation of any problem. If the solution to any problem can be formulated recursively using the solution to its sub-problems, and if its sub-problems are overlapping, then one can easily `memoize or store the solutions to the sub-problems in a table`. Whenever we attempt to solve a new sub-problem, we first check the table to see if it is already solved. If a solution has been recorded, we can use it directly, otherwise we solve the sub-problem and add its solution to the table.
* **Bottom-up approach(Tabulation)**: Once we formulate the solution to a problem recursively as in terms of its sub-problems, we can try reformulating the problem in a bottom-up fashion: `try solving the sub-problems first and use their solutions to build-on and arrive at solutions to bigger sub-problems`. This is also usually done in a tabular form by iteratively generating solutions to bigger and bigger sub-problems by using the solutions to small sub-problems.

## Example Fibonacci

If we write simple recursive solution for `Fibonacci Numbers`, we get exponential time complexity and if we optimize it by storing solutions of subproblems, time complexity reduces to linear.

***Without dynamic programming, usual recursion:***

```go
func nonDPFibonacci(n int) int {
	if n < 2 {
		return n
	}
	return nonDPFibonacci(n-1) + nonDPFibonacci(n-2)
}
```

***Top-down Dynamic programming, Momoization***

```go
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
```

***Bottom-up Dynamic Programming, Tabulation***

```go
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
```
