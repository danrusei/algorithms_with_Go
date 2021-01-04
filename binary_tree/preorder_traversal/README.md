# Check if a given array can represent Preorder Traversal of Binary Search Tree 

Based on: https://www.geeksforgeeks.org/check-if-a-given-array-can-represent-preorder-traversal-of-binary-search-tree/amp/

Given an array of numbers, return true if given array can represent preorder traversal of a Binary Search Tree, else return false. Expected time complexity is O(n).

## Example

**Input**:  {2, 4, 3}  
**Output**: true  
Given array can represent preorder traversal
of below tree  
    2  
     \  
      4  
     /  
    3  

**Input**:  {2, 4, 1}  
**Output**: false  
Given array cannot represent preorder traversal
of a Binary Search Tree.

**Input**:  {40, 30, 35, 80, 100}  
**Output**: true  
Given array can represent preorder traversal
of below tree  
     40  
   /...\  
 30    80  
  \ ......\  
  35     100  

**Input**:  {40, 30, 35, 20, 80, 100}  
**Output**: false  
Given array cannot represent preorder traversal
of a Binary Search Tree.

## Algorithm

Binary Tree Preorder traversal is implemented in the binarytree package. Therefore, here is just used to test out if the list is correct.

```bash
$ go test -v  
=== RUN   TestPreorderTraversal  
=== RUN   TestPreorderTraversal/length_3_true  
Input: [2 4 3]  
Preorder traversal: [2 4 3]  
=== RUN   TestPreorderTraversal/length_3_false  
Input: [2 4 1]  
Preorder traversal: [2 1 4]  
=== RUN   TestPreorderTraversal/length_5_true  
Input: [40 30 35 80 100]  
Preorder traversal: [40 30 35 80 100]  
=== RUN   TestPreorderTraversal/length_6_false  
Input: [40 30 35 20 80 100]  
Preorder traversal: [40 30 20 35 80 100]  
--- PASS: TestPreorderTraversal (0.00s)  
    --- PASS: TestPreorderTraversal/length_3_true (0.00s)  
    --- PASS: TestPreorderTraversal/length_3_false (0.00s)  
    --- PASS: TestPreorderTraversal/length_5_true (0.00s)  
    --- PASS: TestPreorderTraversal/length_6_false (0.00s)  
PASS  
ok  	github.com/danrusei/algorithms_with_go/binary_tree/preorder_traversal	0.001s
```
