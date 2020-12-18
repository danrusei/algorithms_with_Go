# Count number of ways to cover a distance

Source: [GeeksforGeeks](https://www.geeksforgeeks.org/count-number-of-ways-to-cover-a-distance/)

Given a distance ‘dist, count total number of ways to cover the distance with 1, 2 and 3 steps.

## Algorithm

Can be done in 2 ways:

* the first way is to create a recursive structure and store the value in a HashMap and whenever the function is called, return the value store without computing (Top-Down Approach).
* the second way is to take an extra space of size n and start computing values of states from 1, 2 .. to n, i.e. compute values of i, i+1, i+2 and then use them to calculate the value of i+3 (Bottom-Up Approach).

This script is using Bottom-Up approach:

* create an array of size n + 1 and initilize the first 3 variables with 1, 1, 2. The base cases.
* run a loop from 3 to n.
* for each index i, compute the value of i position as dp[i] = dp[i-1] + dp[i-2] + dp[i-3].
* print the value of dp[n], as the Count of number of ways to cover a distance.
