# Detect and Remove Loop in a Linked List

From **[GeeksforGeeks](https://www.geeksforgeeks.org/detect-and-remove-loop-in-a-linked-list/amp/)**

## Example

Given linkedlist:  
1 -> 2 -> 3 -> 4 -> 5  
          |         |  
          8 <- 7 <- 6

Detect and remove the loop.

Resulted linkedlist:  
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8

## Algorithm

* use Floyd’s Cycle detection algorithm that terminates when fast and slow pointers meet at a common point. 
* the common point can be one of the loop nodes (3 or 4 or 5 or 6 or 7 or 8 in the above example)
* store the address of this in a pointer variable say `ptr2`
* start from the head of the Linked List (`ptr1`) and check for nodes one by one if they are reachable from `ptr2`
* whenever we find a node that is reachable, we know that this node is the starting node of the loop in Linked List and we can get the pointer to the previous of this node
* set the `ptr2` to nil to break the loop
