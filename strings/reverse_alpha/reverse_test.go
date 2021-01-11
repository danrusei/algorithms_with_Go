package reverse

import "testing"

func TestReverse(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverse(tc.input)
			if result != tc.output {
				t.Errorf("Expected: %s, got: %s", tc.output, result)
			}

		})
	}
}
