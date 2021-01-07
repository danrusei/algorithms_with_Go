# Check if a given array can represent Preorder Traversal of Binary Search Tree

Based on: https://www.geeksforgeeks.org/bottom-view-binary-tree/amp/

Given a Binary Tree, we need to print the bottom view from left to right. A node x is there in output if x is the bottommost node at its horizontal distance. Horizontal distance of left child of a node x is equal to horizontal distance of x minus 1, and that of right child is horizontal distance of x plus 1.

## Algorithm

* Tree nodes are added in a queue to do pre-order traversal of the tree
* A HashMap  is used to store: key value pairs as **key = column** and **value = a struct** which contain the Value of the Node and the **level in the tree**
* Start with the horizontal distance(col) = 0 and vertical distance (level) = 0 of the root node
* recursively traverse left and right subtree, by adding left child to queue along with the horizontal distance as col-1 and right child as col+1.

## Example

Suppose the binary tree is:

```bash
         1
        /  \
       2    3
      / \  / \
     4   5 6  7
            \
             8
```

Bottom view of the tree is: **4 2 5 6 8 7**

> Horizontal distance of node with value 1: 0, level = 1  
Horizontal distance of node with value 2: 0 - 1 = -1, level = 2  
Horizontal distance of node with value 3: 0 + 1 = 1, level = 2  
Horizontal distance of node with value 4: -1 - 1 = -2, level = 3  
Horizontal distance of node with value 5: -1 + 1 = 0, level = 3  
Horizontal distance of node with value 6: 1 - 1 = 0, level = 3  
Horizontal distance of node with value 7: 1 + 1 = 2, level = 3  
Horizontal distance of node with value 8: 0 + 1 = 1, level = 4  
Print the nodes which appear in the last level of that line.  
for hd = -2, print 4  
for hd = -1, print 2  
for hd = 0, print 5 and 6 because they both appear in the last level of that vertical line  
for hd = 1, print 8  
for hd = 2, print 7  
