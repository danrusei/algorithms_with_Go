# Select a Random Node from a Singly Linked List

Problem description: [GeeksforGeeks](https://www.geeksforgeeks.org/select-a-random-node-from-a-singly-linked-list/amp/)

Given a singly linked list, select a random node from linked list (the probability of picking a node should be 1/N if there are N nodes in list). You are given a random number generator.

## Algorithm

* initialize result as first node
* initialize n = 2
* traverse the linked list
  * generate a random number from 0 to n-1.
  * if the number is equal to 0, then replace result with current node.
  * n++

## Result

```bash
$ go run main.go 
7
$ go run main.go 
1
$ go run main.go 
9
$ go run main.go 
3
$ go run main.go 
10
$ go run main.go 
9
$ go run main.go 
5
```
