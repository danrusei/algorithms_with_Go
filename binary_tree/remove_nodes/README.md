# Remove nodes on root to leaf paths of length < K

Problem description: [GeeksforGeeks](https://www.geeksforgeeks.org/remove-nodes-root-leaf-paths-length-k/amp/)

Given a Binary Tree and a number k, remove all nodes that lie only on root to leaf path(s) of length smaller than k. If a node X lies on multiple root-to-leaf paths and if any of the paths has path length >= k, then X is not deleted from Binary Tree. In other words a node is deleted if all paths going through it have lengths smaller than k.

## Example

Consider the following example Binary Tree

```bash
               1
           /      \
         2          3
      /     \         \
    4         5        6
  /                   /
 7                   8 

Input: Root of above Binary Tree
       k = 4

Output: The tree should be changed to following  
           1
        /     \
      2          3
     /             \
   4                 6
 /                  /
7                  8

There are 3 paths 
i) 1->2->4->7      path length = 4
ii) 1->2->5        path length = 3
iii) 1->3->6->8    path length = 4 

There is only one path " 1->2->5 " of length smaller than 4.  The node 5 is the only node that lies only on this path, so node 5 is removed. Nodes 2 and 1 are not removed as they are parts of other paths of length 4 as well.

If k is 5 or greater than 5, then whole tree is deleted. 
If k is 3 or less than 3, then nothing is deleted.
```

## Algorithm

* Traverse the tree in postorder fashion so that if a leaf node path length is shorter than k, then that node and  all of its descendants till the node which are not on some other path are removed.
* If root is a leaf node and it's level is less than k then remove this node.

## Result

```bash
$ go run main.go 
In order traversal of the original tree:
((((7) 4) 2 (5)) 1 (3 ((8) 6)))
In order traversal of the modified tree:
((((7) 4) 2) 1 (3 ((8) 6)))
```
