package main

import "fmt"

func swapBits(x int) int {

	// Get all even bits of x
	evenBits := x & 0xAAAAAAAA

	// Get all odd bits of x
	oddBits := x & 0x55555555

	// Right shift even bits
	evenBits >>= 1

	// Left shift odd bits
	oddBits <<= 1

	return evenBits | oddBits
}

func main() {

	x := 23
	y := swapBits(x)

	fmt.Printf("Original number is %d and swapped number is %d\n", x, y)
}
