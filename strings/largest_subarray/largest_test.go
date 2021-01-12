package largest

import "testing"

func TestReverse(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := largestSlice(tc.input)
			if result != tc.output {
				t.Errorf("Expected: %d, got: %d", tc.output, result)
			}

		})
	}
}
