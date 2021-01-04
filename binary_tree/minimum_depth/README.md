# Find Minimum Depth of a Binary Tree 

Based on: https://www.geeksforgeeks.org/find-minimum-depth-of-a-binary-tree/amp/

Given a binary tree, find its minimum depth. The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

## Example

Given the  below Binary Tree:  
![binary tree](tree122.gif)

Minimum height of below Binary Tree is 2. Note that the path must end on a leaf node.

## Algorithm

* traverse the given Binary Tree and for every node, check if it is a leaf node
* If yes, then return 1
* If not leaf node then if the left subtree is NULL, then recur for the right subtree
* if the right subtree is NULL, then recur for the left subtree
* if both left and right subtrees are not NULL, then take the minimum of two heights.
