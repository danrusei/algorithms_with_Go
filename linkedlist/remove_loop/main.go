package main

import (
	"fmt"

	"github.com/danrusei/algorithms_with_go/linkedlist"
)

type llist struct {
	linkedlist.Linkedlist
}

//DetectAndRemove detects the loop, if any returns true and runs removeLoop function to remove the loop
func (ll *llist) DetectAndRemove() bool {
	slow := ll.Head
	fast := ll.Head

	//for loop is running until it reach the end of the list, therefore there is no loop
	//or the loop has been identified as the `slow` node is equal to `fast` node
	for {
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			ll.removeLoop(slow)
			return true
		}
	}

	return false
}

// loopNode --> Pointer to one of the loop nodes
func (ll *llist) removeLoop(loopNode *linkedlist.Node) {

	ptr1 := ll.Head
	ptr2 := &linkedlist.Node{}

	//we know that the linkedlist has a loop
	//ptr2 is an element inside the loop, can be 3, 4, 5, 6, 7 or 8
	//ptr1 advance node by node until ptr2.Next == ptr1
	for {
		ptr2 = loopNode
		for ; ptr2.Next != loopNode && ptr2.Next != ptr1; ptr2 = ptr2.Next {
		}
		if ptr2.Next == ptr1 {
			break
		}

		ptr1 = ptr1.Next
	}

	// remove the loop by seting the pointer of last element to nil
	ptr2.Next = nil
}

func main() {
	list := llist{}
	list.AddAtEnd(1)
	list.AddAtEnd(2)
	list.AddAtEnd(3)
	list.AddAtEnd(4)
	list.AddAtEnd(5)
	list.AddAtEnd(6)
	list.AddAtEnd(7)
	list.AddAtEnd(8)

	endNode := list.Head
	for ; endNode.Next != nil; endNode = endNode.Next {
	}
	//create the loop
	endNode.Next = list.Head.Next.Next

	fmt.Println(list.DetectAndRemove())
	fmt.Println(list.String())
}
