package main

import "fmt"

// recursively call this function untill all bits are shifted
func bin(x int) {

	if x > 1 {
		bin(x >> 1)
	}
	// print last bit from the remainder
	fmt.Printf("%d", x&1)
}

func main() {
	x := 131
	fmt.Printf("Using recursive function, the binary of %d is: ", x)
	bin(x)
	fmt.Println()
	//builtin using fmt package
	fmt.Printf("Using fmt package, the binary of %d is: %b\n", x, x)
}
