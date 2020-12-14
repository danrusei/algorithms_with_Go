package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mergesort(tc.original)
			if !reflect.DeepEqual(tc.original, tc.sorted) {
				t.Errorf("got: %v, want: %v", tc.original, tc.sorted)
			}
		})
	}
}
