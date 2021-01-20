# Rotate bits of a number 

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/rotate-bits-of-an-integer/amp/)  

Bit Rotation: A rotation (or circular shift) is an operation similar to shift except that the bits that fall off at one end are put back to the other end.

* In left rotation, the bits that fall off at left end are put back at right end.
* In right rotation, the bits that fall off at right end are put back at left end.

## Example

Let n = 3 is stored using 8 bits.

* Left rotation of n = 11100101 by 3 makes n = 00101111
* Right rotation of n = 11100101 by 3 makes n = 10111100

## Algorithm

Left Rotation:

* n << rotate , leaves last "rotate" bits 0
* to populate these with the first "rotate" bits of the original number
* move first bits of the number to the end (n >> (INT_BITS - rotate))
* and do bitwise OR with the rotated number

Right Rotation:

* n >> rotate , leaves first "rotate" bits 0
* to populate these with the last "rotate" bits of the original number
* move last bits of the number to the begining (n << (INT_BITS - rotate))
* and do bitwise OR with the rotated number

## Result

```bash
$ go run main.go 
The number 152 in bits: 10011000
After left rotation the number is 196 and in bits: 11000100
After right rotation the number is 19 and in bits: 10011
```
