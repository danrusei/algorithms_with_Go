# Convert array into Zig-Zag fashion

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/convert-array-into-zig-zag-fashion/amp/)

Given an array of DISTINCT elements, rearrange the elements of array in zig-zag fashion in O(n) time. The converted array should be in form a < b > c < d > e < f.

## Example

```bash
Input: arr[] = {4, 3, 7, 8, 6, 2, 1} 
Output: arr[] = {3, 7, 4, 8, 2, 6, 1}

Input: arr[] = {1, 4, 3, 2} 
Output: arr[] = {1, 4, 2, 3}
```

## Algorithm

* create a flag to denote the "<" or ">"
* if two consecutive numbers do not have the expected relationship, then swap the elements

## Result

```bash
$ go test -v
=== RUN   TestReverse
=== RUN   TestReverse/first_10
=== RUN   TestReverse/7_values
=== RUN   TestReverse/4_values
--- PASS: TestReverse (0.00s)
    --- PASS: TestReverse/first_10 (0.00s)
    --- PASS: TestReverse/7_values (0.00s)
    --- PASS: TestReverse/4_values (0.00s)
PASS
```
