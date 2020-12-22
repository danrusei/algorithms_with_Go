# Subset Sum Problem 

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/count-number-of-ways-to-cover-a-distance/)

Given a set of non-negative integers, and a value sum, determine if there is a subset of the given set with sum equal to given sum. 

## Example

> Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 9  
> Output: True  
> There is a subset (4, 5) with sum 9.
>
> Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 30  
> Output: False  
> There is no subset that add up to 30.

## Algorithm

* bottom-up approach computes `mat[i][j]` for each `1 <= i <= n` and `1 <= j <= sum`
* is `true` if subset with sum `j` can be found using items up to first `i` items
* it uses values of smaller values `i` and `j` already computed
