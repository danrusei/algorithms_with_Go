package hashtable

import "testing"

func TestAddKeyValue(t *testing.T) {
	var h *HashTable
	h = New()

	testcases := []struct {
		key   string
		value int
	}{
		{"first", 1},
		{"second", 2},
		{"third", 3},
		{"fourth", 4},
		{"fifth", 5},
		{"sixth", 6},
		{"seventh", 7},
		{"eigth", 8},
		{"nineth", 9},
		{"tenth", 10},
	}
	for _, tc := range testcases {
		h.AddKeyValue(tc.key, tc.value)
		result, ok := h.GetValueForKey(tc.key)
		if !ok {
			t.Errorf("Could not get the value for key: %s", tc.key)
		}
		if result.(int) != tc.value {
			t.Errorf("Wrong value for key: %s, got: %d, expected: %d", tc.key, result, tc.value)
		}
	}
}
