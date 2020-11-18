package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	for _, tc := range testcases {
		tc.queue.Enqueue(tc.element)

		if len(tc.queue) != len(tc.enqueue) {
			t.Fatalf("Fail, the lengths between current queue and expected are different")
		}

		for i := range tc.queue {
			if tc.queue[i] != tc.enqueue[i] {
				t.Fatalf("Fail, expected: %v, got: %v", tc.enqueue, tc.queue)
			}
		}
	}
}

func TestDequeue(t *testing.T) {
	for _, tc := range testcases {
		element, ok := tc.queue.Dequeue()
		if ok != tc.ok {
			t.Fatalf("Fail, got: %v, expected: %v", ok, tc.ok)
		}

		if element != tc.dequeue {
			t.Fatalf("Fail got: %d, expected: %d", element, tc.dequeue)
		}
	}
}
