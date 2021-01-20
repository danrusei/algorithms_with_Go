package main

import "fmt"

func nextSparse(x int) int {

	// bin store the binary representation of x
	// bin[0] contains least significant bit
	bin := []int{}
	for x != 0 {
		bin = append(bin, x&1)
		x >>= 1
	}

	last := 0

	for i := 1; i < len(bin); i++ {

		if bin[i] == 1 && bin[i-1] == 1 && bin[i+1] != 1 {

			// make the next bit 1
			bin[i+1] = 1

			// make all bits before current bit as 0 to make
			// sure that we get the smallest next number
			for j := i; j >= last; j-- {
				bin[j] = 0

			}

			// store position of the bit set so that this bit
			// and bits before it are not changed next time.
			last = i + 1
		}
	}

	//create and return the decimal equivalent of modified bin[]
	result := 0
	for i := 0; i < len(bin); i++ {
		result += bin[i] * (1 << i)
	}

	return result
}

func main() {
	x := 38
	y := nextSparse(x)

	fmt.Printf("next sparse number starting from %d is %d\n", x, y)

}
