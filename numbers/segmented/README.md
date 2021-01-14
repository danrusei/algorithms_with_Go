# Segmented Sieve 

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/segmented-sieve/amp/)  
Source: [Wikipedia - Segmented Sieves](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Segmented_sieve)

Given a number n, print all primes smaller than or equal to n.

The implemented [Sieve of Eratosthenes](https://github.com/danrusei/algorithms_with_Go/tree/main/numbers/eratosthenes) looks good, but consider the situation when n is large, the Simple Sieve faces following issues:

* an array of size Θ(n) may not fit in memory
* the simple Sieve is not cache friendly even for slightly bigger n. The algorithm traverses the array without locality of reference

## Example

> Input : n =10  
> Output : 2 3 5 7
>
> Input : n = 20  
> Output: 2 3 5 7 11 13 17 19

## Algorithm

The idea of segmented sieve is to divide the range [0..n-1] in different segments and compute primes in all segments one by one. This algorithm first uses Simple Sieve to find primes smaller than or equal to √(n).

* use Simple Sieve to find all primes upto square root of ‘n’ and store these primes in an array “prime[]”. Store the found primes in an array ‘prime[]’.
* divide [0..n-1] range in different segments such that size of every segment is at-most √n
* for every segment [low..high]
  * create an array mark[high-low+1]. Here we need only O(x) space where x is number of elements in given range.
  * iterate through all primes found in step 1. For every prime, mark its multiples in given range [low..high].

In Simple Sieve, we needed O(n) space which may not be feasible for large n. Here we need O(√n) space and we process smaller ranges at a time (locality of reference)

## Result

```bash
 go test -v
=== RUN   TestPrimes
=== RUN   TestPrimes/10_value
=== RUN   TestPrimes/20_value
=== RUN   TestPrimes/99_value
--- PASS: TestPrimes (0.00s)
    --- PASS: TestPrimes/10_value (0.00s)
    --- PASS: TestPrimes/20_value (0.00s)
    --- PASS: TestPrimes/99_value (0.00s)
PASS
```
