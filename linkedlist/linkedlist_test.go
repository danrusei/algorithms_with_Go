package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedlist()
	list.AddAtBeg(1)
	list.AddAtBeg(2)
	list.AddAtBeg(3)

	t.Run("Test AddAtBeg()", func(t *testing.T) {
		want := []int{3, 2, 1}
		got := []int{}
		current := list.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list.AddAtEnd(4)

	t.Run("Test AddAtEnd()", func(t *testing.T) {
		want := []int{3, 2, 1, 4}
		got := []int{}
		current := list.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtBeg()", func(t *testing.T) {
		want := int(3)
		got := list.DelAtBeg()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtEnd()", func(t *testing.T) {
		want := int(4)
		got := list.DelAtEnd()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DeleteWithValute", func(t *testing.T) {
		want := 1
		got := list.DeleteWithValute(1)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
