# Reverse alternate levels of a perfect binary tree

Problem description: [GeeksforGeeks](https://www.geeksforgeeks.org/reverse-alternate-levels-binary-tree/amp/)

Given a Perfect Binary Tree, reverse the alternate level nodes of the binary tree. 

## Example

```bash
Given tree: 
               11
            /     \
          12       13
         /  \     /  \
       14    15   16  17
       / \  / \  / \  / \
      1   2 3  4 5  6 7  8 

Modified tree:
              11
            /     \
          13       12
         /  \     /  \
       14    15  16   17
       / \  / \  / \  / \
      8  7 6  5 4  3  2  1 


```

## Algorithm

* traverse the given tree in inorder fashion and store all odd level values in a slice
* reverse the slice
* traverse the tree again inorder fashion, take elements from the slice and store the values to every odd level traversed node.

## Result

```bash
$ go run main.go 
In order traversal of the original tree:
((((1) 14 (2)) 12 ((3) 15 (4))) 11 (((5) 16 (6)) 13 ((7) 16 (8))))
In order traversal of the modified tree:
((((8) 14 (7)) 13 ((6) 15 (5))) 11 (((4) 16 (3)) 12 ((2) 16 (1))))
```
