# Given a linked list which is sorted, how will you insert in sorted way

Based on: https://www.geeksforgeeks.org/given-a-linked-list-which-is-sorted-how-will-you-insert-in-sorted-way/

## Example

Given linkedlist:  
2 -> 5 -> 7 -> 10 -> 15

Insert value 9

Resulted linkedlist:  
2 -> 5 -> 7 -> 9 -> 10 -> 15

## Algorithm

* If Linked list is empty then make the node as head and return it.
* If the value of the node to be inserted is smaller than the value of the head node, then insert the node at the start and make it head.
* In a loop, find the appropriate node after which the input node (let 9) is to be inserted. To find the appropriate node start from the head, keep moving until you reach a node GN (10 in the below diagram) who's value is greater than the input node. The node just before GN is the appropriate node (7).
* Insert the node (9) after the appropriate node (7) found in step 3.
