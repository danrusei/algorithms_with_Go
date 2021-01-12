# Reverse a string without affecting special characters

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/reverse-a-string-without-affecting-special-characters/amp/)

Given a string, that contains special character together with alphabets (‘a’ to ‘z’ and ‘A’ to ‘Z’), reverse the string in a way that special characters are not affected.

## Example

> Input:   str = "a,b$c"  
> Output:  str = "c,b$a"  
> Note that $ and , are not moved anywhere.  
> Only subsequence "abc" is reversed
>
> Input:   str = "Ab,c,de!$"
> Output:  str = "ed,c,bA!$"

## Algorithm

* Let input string be `str` and length of string be `n`
* l = 0, r = n-1
* for `l` is smaller than `r`, do following
  * If `str[l]` is not an alphabetic character, do `l++`
  * else if `str[r]` is not an alphabetic character, do `r--`
  * else swap `str[l]` and `str[r]`

## Result

```bash
$ go test -v
=== RUN   TestReverse
=== RUN   TestReverse/simple
=== RUN   TestReverse/with_punctuation
=== RUN   TestReverse/with_capital
=== RUN   TestReverse/with_trailing_space
--- PASS: TestReverse (0.00s)
    --- PASS: TestReverse/simple (0.00s)
    --- PASS: TestReverse/with_punctuation (0.00s)
    --- PASS: TestReverse/with_capital (0.00s)
    --- PASS: TestReverse/with_trailing_space (0.00s)
PASS
ok  	_/home/rdan/projects_public/algorithms_with_Go/strings/reverse_alpha	0.001s
```
