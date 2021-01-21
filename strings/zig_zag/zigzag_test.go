package zigzag

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			zigzag(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("Expected: %d, got: %d", tc.output, tc.input)
			}
		})
	}
}
