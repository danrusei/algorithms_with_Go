# Find Next Sparse Number 

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/given-a-number-find-next-sparse-number/amp/)  

Given a number x, find the smallest Sparse number which greater than or equal to x.  
A number is Sparse if there are no two adjacent 1s in its binary representation. For example 5 (binary representation: 101) is sparse, but 6 (binary representation: 110) is not sparse.

## Example

```bash
Input: x = 6
Output: Next Sparse Number is 8

Input: x = 4
Output: Next Sparse Number is 4

Input: x = 38
Output: Next Sparse Number is 40

Input: x = 44
Output: Next Sparse Number is 64
```

## Algorithm

There is a simple solution which check if the number is sparse, if not it moves to the next decimal number and so on.  
Although this is a valid solution is not so efficient, a more efficient solution is described below:

* store the binary of a given number in a slice  
* initialize the last finalized bit position as 0
* traverse the slice, and start with least significatn bit (from right to left)
  * if we find two adjacent 1's such that next (or third) bit is not 1, then:
    * Make all bits after this 1 to last finalized bit 0
    * Update last finalized bit as next bit

## Result

```bash
$ go run main.go 
next sparse number starting from 38 is 40
```
