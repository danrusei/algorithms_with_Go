package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/danrusei/algorithms_with_go/linkedlist"
)

type llist struct {
	linkedlist.Linkedlist
}

// we can easily select a random node by using the length field of the linkedlist
// let's assume this time that the length field has not been defined
func (ll *llist) randomNode() int {
	if ll.Head == nil {
		return 0
	}

	//NewSource returns a new pseudo-random Source seeded with the given value.
	r1 := rand.NewSource(time.Now().UnixNano())
	//New returns a new Rand that uses random values from x1 to generate other random values.
	r2 := rand.New(r1)

	result := ll.Head.Val
	cur := ll.Head

	// traverse the list and increase the n, which is the length of the yet discovered list
	n := 2
	for ; cur.Next != nil; cur = cur.Next {

		if r2.Int()%n == 0 {
			result = cur.Val
		}
		n++
	}

	// type assertion is required because Linked List value = interface{}
	return result.(int)
}

func main() {
	list := llist{}
	list.AddAtEnd(5)
	list.AddAtEnd(8)
	list.AddAtEnd(10)
	list.AddAtEnd(3)
	list.AddAtEnd(1)
	list.AddAtEnd(9)
	list.AddAtEnd(7)
	list.AddAtEnd(2)

	result := list.randomNode()
	fmt.Println(result)
}
