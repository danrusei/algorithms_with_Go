package reverse

import (
	"fmt"
	"unicode"
)

func reverse(input string) string {
	toSlice := make([]rune, len(input))
	for i, c := range input {
		toSlice[i] = c
	}
	reverseSlice(toSlice)

	return fmt.Sprint(string(toSlice))
}

func reverseSlice(input []rune) {
	l := 0
	r := len(input) - 1

	for l < r {
		if !unicode.IsLetter(input[l]) {
			l++
		} else if !unicode.IsLetter(input[r]) {
			r--
		} else {
			input[l], input[r] = input[r], input[l]
			l++
			r--
		}
	}
}
