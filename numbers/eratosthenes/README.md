# Sieve of Eratosthenes

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/sieve-of-eratosthenes/amp/)  
Source: [Wikipedia - Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)

Given a number n, print all primes smaller than or equal to n. It is also given that n is a small number.

## Example

> Input : n =10  
> Output : 2 3 5 7
>
> Input : n = 20  
> Output: 2 3 5 7 11 13 17 19

## Algorithm

A prime number is a natural number that has exactly two distinct natural number divisors: the number 1 and itself.

To find all the prime numbers less than or equal to a given integer n by Eratosthenes' method:

* Create a list of consecutive integers from 2 through n: (2, 3, 4, ..., n).
* Initially, let p equal 2, the smallest prime number.
* Enumerate the multiples of p by counting in increments of p from 2p to n, and mark them in the list (these will be 2p, 3p, 4p, ...; the p itself should not be marked).
* Find the smallest number in the list greater than p that is not marked. If there was no such number, stop. Otherwise, let p now equal this new number (which is the next prime), and repeat from step 3.
* When the algorithm terminates, the numbers remaining not marked in the list are all the primes below n.

The main idea here is that every value given to p will be prime, because if it were composite it would be marked as a multiple of some other, smaller prime. Note that some of the numbers may be marked more than once (e.g., 15 will be marked both for 3 and 5).

As a refinement, it is sufficient to mark the numbers **in step 3 starting from p * p**, as all the smaller multiples of p will have already been marked at that point. This means that the algorithm is allowed to terminate in step 4 when p2 is greater than n.

## Result

```bash
$ go test -v
=== RUN   TestReverse
=== RUN   TestReverse/10_value
=== RUN   TestReverse/20_value
=== RUN   TestReverse/99_value
--- PASS: TestReverse (0.00s)
    --- PASS: TestReverse/10_value (0.00s)
    --- PASS: TestReverse/20_value (0.00s)
    --- PASS: TestReverse/99_value (0.00s)
PASS
```
