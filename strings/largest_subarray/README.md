# Length of the largest subarray with contiguous elements 

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/length-largest-subarray-contiguous-elements-set-1/amp/)

Given an array of distinct integers, find length of the longest subarray which contains numbers that can be arranged in a continuous sequence.

## Example

> Input:  arr[] = {10, 12, 11};  
> Output: Length of the longest contiguous subarray is 3
>
> Input:  arr[] = {14, 12, 11, 20};  
> Output: Length of the longest contiguous subarray is 2
>
> Input:  arr[] = {1, 56, 58, 57, 90, 92, 94, 93, 91, 45};  
> Output: Length of the longest contiguous subarray is 5

## Algorithm

If all elements are distinct, then a subarray has contiguous elements if and only if the difference between maximum and minimum elements in subarray is equal to the difference between last and first indexes of subarray. So the idea is to keep track of minimum and maximum element in every subarray.

## Result

```bash
$ go test -v
=== RUN   TestReverse
=== RUN   TestReverse/entire_slice
=== RUN   TestReverse/middle
=== RUN   TestReverse/larger_slice
--- PASS: TestReverse (0.00s)
    --- PASS: TestReverse/entire_slice (0.00s)
    --- PASS: TestReverse/middle (0.00s)
    --- PASS: TestReverse/larger_slice (0.00s)
PASS
ok  	_/home/rdan/projects_public/algorithms_with_Go/strings/largest_subarray	0.001s
```
