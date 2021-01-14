package primes

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := primes(tc.num)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %d, got: %d", tc.expected, result)
			}
		})
	}
}
