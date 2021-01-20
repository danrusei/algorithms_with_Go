package main

import "fmt"

// INT_BITS the number n is stored using 8 bits
const INT_BITS = 8

// function to left rotate n by rotate bits
func rotateLeft(n int, rotate int) uint8 {
	// n << rotate , leaves last "rotate" bits 0
	// to populate these with the first "rotate" bits of the original number
	// move first bits of the number to the end (n >> (INT_BITS - rotate))
	// and do bitwise OR with the rotated number
	return uint8((n << rotate) | (n >> (INT_BITS - rotate)))
}

// function to right rotate n by rotate bits
func rotateRight(n int, rotate int) uint8 {
	// n >> rotate , leaves first "rotate" bits 0
	// to populate these with the last "rotate" bits of the original number
	// move last bits of the number to the begining (n << (INT_BITS - rotate))
	// and do bitwise OR with the rotated number
	return uint8((n >> rotate) | (n << (INT_BITS - rotate)))
}

func main() {
	n := 152
	rotate := 3

	fmt.Printf("The number %d in bits: %b\n", n, n)

	rl := rotateLeft(n, rotate)
	fmt.Printf("After left rotation the number is %d and in bits: %b\n", rl, rl)

	rr := rotateRight(n, rotate)
	fmt.Printf("After right rotation the number is %d and in bits: %b\n", rr, rr)
}
