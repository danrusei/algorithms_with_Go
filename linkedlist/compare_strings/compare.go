package compare

func compareStrings(list1 *llist, list2 *llist) int {
	// remember that the list is defines as:
	// type Linkedlist struct {
	//  	Length int
	//	  	Head   *Node
	//  }
	// so we have the length to check initialy without comparing all elements in the lists
	// however if the lists have the same length we have to compare the elements

	if list1.Length > list2.Length {
		return 1
	} else if list1.Length < list2.Length {
		return -1
	}

	cur1 := list1.Head
	cur2 := list2.Head
	// type assertion is required as the Node value = interface{}
	for cur1.Val.(string) == cur2.Val.(string) && cur1.Next != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	if cur1.Val.(string) > cur2.Val.(string) {
		return 1
	} else if cur1.Val.(string) < cur2.Val.(string) {
		return -1
	}
	//asumes the list are equal
	return 0
}
