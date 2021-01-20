# Compare two strings represented as linked lists 

Problem description: [GeeksforGeeks](https://www.geeksforgeeks.org/compare-two-strings-represented-as-linked-lists/amp/)

Given two strings, represented as linked lists (every character is a node in a linked list). Write a function compare() that works similar to strcmp(), i.e., it returns 0 if both strings are same, 1 if first linked list is lexicographically greater, and -1 if the second string is lexicographically greater.

## Example

```bash
Input: list1 = g->e->e->k->s->a
       list2 = g->e->e->k->s->b
Output: -1

Input: list1 = g->e->e->k->s->a
       list2 = g->e->e->k->s
Output: 1

Input: list1 = g->e->e->k->s
       list2 = g->e->e->k->s
Output: 0
```

## Algorithm

* define the linkedlist as below:

```go
type Node struct {
	Val  string
	Next *Node
}

type Linkedlist struct {
	Length int
	Head   *Node
}
```

* compare the length of the 2 lists, and return immediately if they are not the same length
* otherwise compare each element and stop when the lists values are different

## Result

```bash
$ go test -v
=== RUN   TestComparedStrings
=== RUN   TestComparedStrings/second_list_greater
=== RUN   TestComparedStrings/the_same_list
=== RUN   TestComparedStrings/first_list_greater
--- PASS: TestComparedStrings (0.00s)
    --- PASS: TestComparedStrings/second_list_greater (0.00s)
    --- PASS: TestComparedStrings/the_same_list (0.00s)
    --- PASS: TestComparedStrings/first_list_greater (0.00s)
PASS
ok  	github.com/danrusei/algorithms_with_go/linkedlist/compare_strings	0.001s
```
