# Binary representation of a given number

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/binary-representation-of-a-given-number/amp/)
Write a program to print Binary representation of a given number.

## Algorithm

Recursive using bitwise operator

* Check n > 1
* **right shift** the number by 1 bit and recursive call the function
* Print the bits of number

Alternate solution is to use fmt Go package to transform decimal to binary:

```go
fmt.Printf("The binary of the %d is: %b\n", x, x)
```
