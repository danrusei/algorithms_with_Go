package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubblesort(t *testing.T) {
	for _, tc := range testcases {
		got := tc.sort(tc.original)

		if !reflect.DeepEqual(got, tc.sorted) {
			t.Errorf("got: %v, want: %v", got, tc.sorted)

		}
	}
}
