package subsetsum

import "testing"

func TestSubsetSum(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			output := subsetsum(tc.input, tc.sum)
			if output != tc.output {
				t.Errorf("got: %v, want: %v", output, tc.output)
			}

		})
	}
}
