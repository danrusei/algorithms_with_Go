# Smallest subarray with sum greater than a given value

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/minimum-length-subarray-sum-greater-given-value/amp/)

Given an array of integers and a number x, find the smallest subarray with sum greater than the given value.

## Example

> arr[] = {1, 4, 45, 6, 0, 19}  
>    x  =  51  
> Output: 3  
> Minimum length subarray is {4, 45, 6}
>
> arr[] = {1, 10, 5, 2, 7}  
>    x  = 9  
> Output: 1  
> Minimum length subarray is {10}
>
> arr[] = {1, 11, 100, 1, 0, 200, 3, 2, 1, 250}  
>     x = 280  
> Output: 4  
> Minimum length subarray is {100, 1, 0, 200}  
>
> arr[] = {1, 2, 4}  
>     x = 8  
> Output : Not Possible  
> Whole array sum is smaller than 8.

## Algorithm

Use two nested loops:

* the outer loop picks a starting element
* the inner loop considers all elements from the right side of the current start
* Whenever sum of elements between current start and end becomes more than the given number, update the result if current length is smaller than the smallest length so far

## Result

```bash
$ go test -v
=== RUN   TestReverse
=== RUN   TestReverse/51_value
=== RUN   TestReverse/9_value
=== RUN   TestReverse/280_value
=== RUN   TestReverse/8_value
--- PASS: TestReverse (0.00s)
    --- PASS: TestReverse/51_value (0.00s)
    --- PASS: TestReverse/9_value (0.00s)
    --- PASS: TestReverse/280_value (0.00s)
    --- PASS: TestReverse/8_value (0.00s)
PASS
ok  	_/home/rdan/projects_public/algorithms_with_Go/strings/smallest_subarray	0.001s
```
