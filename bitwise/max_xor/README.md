# Find the maximum subarray XOR in a given array

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/find-the-maximum-subarray-xor-in-a-given-array/amp/)  
Given an array of integers. find the maximum XOR subarray value in given array. Expected time complexity O(n).

> Input: arr[] = {1, 2, 3, 4}  
Output: 7  
The subarray {3, 4} has maximum XOR value
>
> Input: arr[] = {8, 1, 2, 12, 7, 6}  
Output: 15  
The subarray {1, 2, 12} has maximum XOR value
>
> Input: arr[] = {4, 6}  
Output: 6  
The subarray {6} has maximum XOR value

## Solution

* Iterate 2 times over the slice of integers, second time only the ending part
* `currXor` is changed iteratively as it progress through the loop
* `response` value is updated every time when the computed XOR is bigger than the current value
