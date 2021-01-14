package primes

import (
	"reflect"
	"testing"
)

func TestPrimes(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := segmentedSieve(tc.num)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %d, got: %d", tc.expected, result)
			}
		})
	}
}
