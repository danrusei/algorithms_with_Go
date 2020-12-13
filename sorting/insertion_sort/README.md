# Insertion Sort

Source: [Wikipedia](https://en.wikipedia.org/wiki/Insertion_sort)

Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time.

## Complexity

Worst Case Time Complexity [ Big-O ]: **O(n2)**  
Best Case Time Complexity [Big-omega]: **O(n)**  
Average Time Complexity [Big-theta]: **O(n2)**  
Space Complexity: **O(n)**

## Algorithm

Sorting is typically done in-place, by iterating up the array, growing the sorted list behind it.

* At each array-position, it checks the value there against the largest value in the sorted list (which happens to be next to it, in the previous array-position checked).
* If larger, it leaves the element in place and moves to the next.
* If smaller, it finds the correct position within the sorted list, shifts all the larger values up to make a space, and inserts into that correct position.

## Example

> **12**, 11, 13, 5, 6
>
> Let us loop for i = 1 (second element of the array) to 4 (last element of the array)  
> i = 1. Since 11 is smaller than 12, move 12 and insert 11 before 12  
> **11**, 12, 13, 5, 6
>
> i = 2. 13 will remain at its position as all elements in A[0..I-1] are smaller than 13  
> **11, 12, 13**, 5, 6
>
> i = 3. 5 will move to the beginning and all other elements from 11 to 13 will move one position ahead of their current position.  
> **5, 11, 12, 13**, 6
>
> i = 4. 6 will move to position after 5, and elements from 11 to 13 will move one position ahead of their current position.  
> **5, 6, 11, 12, 13**
