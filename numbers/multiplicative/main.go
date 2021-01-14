package main

import "fmt"

func modInverse(a int, m int) int {
	m0 := m
	x, y := 1, 0

	if m == 1 {
		return 0
	}

	for a > 1 {
		// q is quotient
		q := a / m
		t := m

		// m is remainder now, process same as Euclid's algo
		m = a % m
		a = t
		t = y

		// update y and x
		y = x - q*y
		x = t
	}

	// make x positive
	if x < 0 {
		x += m0
	}

	return x
}

func main() {
	a, m := 3, 11
	fmt.Printf("Modular multiplicative inverse is %d\n", modInverse(a, m))
}
