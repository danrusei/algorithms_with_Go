# Longest Common Subsequence

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/)  
Source: [Wikipedia](https://en.wikipedia.org/wiki/Longest_common_subsequence_problem)  
Time Complexity: **O(m*n)**

`LCS Problem Statement`: Given two sequences, find the length of longest subsequence present in both of them. A subsequence is a sequence that appears in the same relative order, but not necessarily contiguous. For example, “abc”, “abg”, “bdf”, “aeg”, ‘”acefg”, .. etc are subsequences of “abcdefg”.

LCS for input Sequences “ABCDGH” and “AEDFHR” is “ADH” of length 3.  
LCS for input Sequences “AGGTAB” and “GXTXAYB” is “GTAB” of length 4.

## Algorithm

* The recursive algorithm controls what order we fill in the computed entries in the array L, and we get the same results if it's filled in some other order.
* use a simple nested loop, that visits the array systematically
* when we fill in a cell L[i,j], we need to already know the values it depends on, in this case L[i+1,j], L[i,j+1], and L[i+1,j+1]. For this reason the array is traversed backwards, from the last row working up to the first and from the last column working up to the first.
* this is iterative (because it uses nested loops instead of recursion) or bottom up (because the order we fill in the array is from smaller simpler subproblems to bigger more complicated ones).
