package smallest

import "testing"

func TestReverse(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := smallestSlice(tc.num, tc.input)
			if result == len(tc.input)+1 {
				result = 0
			}
			if result != tc.output {
				t.Errorf("Expected: %d, got: %d", tc.output, result)
			}
		})
	}
}
