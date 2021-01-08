# Bitwise Operation

In computer programming, a bitwise operation operates on a bit string, a bit array or a binary numeral (considered as a bit string) at the level of its individual bits. It is a fast and simple action, basic to the higher level arithmetic operations and directly supported by the processor. Most bitwise operations are presented as two-operand instructions where the result replaces one of the input operands.

On simple low-cost processors, typically, bitwise operations are substantially faster than division, several times faster than multiplication, and sometimes significantly faster than addition.[clarification needed] While modern processors usually perform addition and multiplication just as fast as bitwise operations due to their longer instruction pipelines and other architectural design choices, bitwise operations do commonly use less power because of the reduced use of resources.

***Source: [Wikipedia](https://en.wikipedia.org/wiki/Bitwise_operation)***

## Go built-in operators

| Operation | Result | Description
| ----------- | ----------- | ----------- |
|0011 & 0101 | 0001 | Bitwise AND
|0011 \| 0101 | 0111 | Bitwise OR
|0011 ^ 0101 | 0110 | Bitwise XOR
|^0101 | 1010 | Bitwise NOT (same as 1111 ^ 0101)
|0011 &^ 0101i | 0010 | Bitclear (AND NOT)
|00110101 << 2 | 11010100 | Left shift
|00110101 << 100 | 00000000 | No upper limit on shift count
|00110101 >> 2 | 00001101 | Right shift
