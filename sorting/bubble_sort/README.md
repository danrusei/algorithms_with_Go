# Bubble Sort

Source: [Wikipedia](https://en.wikipedia.org/wiki/Bubble_sort)

Bubble sort, sometimes referred to as sinking sort, is a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted. The algorithm, which is a comparison sort, is named for the way smaller or larger elements "bubble" to the top of the list.

## Complexity

Worst Case Time Complexity [ Big-O ]: **O(n2)**  
Best Case Time Complexity [Big-omega]: **O(n)**- when the list is already sorted  
Average Time Complexity [Big-theta]: **O(n2)**  
Space Complexity: **O(1)**- because only a single additional memory space is required for temp variable  

## Example

> Take an array of numbers " 5 1 4 2 8", and sort the array from lowest number to greatest number using bubble sort. In each step, elements written in **bold** are being compared. Three passes will be required;
>
> **First Pass**  
> ( **5 1** 4 2 8 ) → ( **1 5** 4 2 8 ), Here, algorithm compares the first two elements, and swaps since 5 > 1.  
> ( 1 **5 4** 2 8 ) → ( 1 **4 5** 2 8 ), Swap since 5 > 4  
> ( 1 4 **5 2**  8 ) → ( 1 4 **2 5** 8 ), Swap since 5 > 2  
> ( 1 4 2 **5 8** ) → ( 1 4 2 **5 8** ), Now, since these elements are already in order (8 > 5), algorithm does not swap them.  
>
> **Second Pass**
> ( **1 4** 2 5 8 ) → ( **1 4** 2 5 8 )  
> ( 1 **4 2** 5 8 ) → ( 1 **2 4** 5 8 ), Swap since 4 > 2  
> ( 1 2 **4 5** 8 ) → ( 1 2 **4 5** 8 )  
> ( 1 2 4 **5 8** ) → ( 1 2 4 **5 8** )  
> Now, the array is already sorted, but the algorithm does not know if it is completed. The algorithm needs one whole pass without any swap to know it is sorted.
>
> **Third Pass**  
> ( **1 2** 4 5 8 ) → ( **1 2** 4 5 8 )  
> ( 1 **2 4** 5 8 ) → ( 1 **2 4** 5 8 )  
> ( 1 2 **4 5** 8 ) → ( 1 2 **4 5** 8 )  
> ( 1 2 4 **5 8** ) → ( 1 2 4 **5 8** )
