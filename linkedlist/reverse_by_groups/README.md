# Reverse a Linked List in groups of given size 

From **[GeeksforGeeks](https://www.geeksforgeeks.org/reverse-a-list-in-groups-of-given-size/amp/)**

## Example

Given a linked list, write a function to reverse every k nodes (where k is an input to the function).

> Example:
>
> Input: 1->2->3->4->5->6->7->8->NULL, K = 3
> Output: 3->2->1->6->5->4->8->7->NULL
>
> Input: 1->2->3->4->5->6->7->8->NULL, K = 5
> Output: 5->4->3->2->1->8->7->6->NULL

## Algorithm

* Reverse the first sub-list of size k. While reversing keep track of the next node and previous node. Let the pointer to the next node be next and pointer to the previous node be prev.
* head->next = reverse(next, k) ( Recursively call for rest of the list and link the two sub-lists )
* Return prev ( prev becomes the new head of the list)
