# Heap Sort

Heap sort is a comparison based sorting technique based on Binary Heap data structure. It is similar to selection sort where we first find the maximum element and place the maximum element at the end. We repeat the same process for the remaining elements.

Source: [Geeksforgeeks](https://www.geeksforgeeks.org/heap-sort/)

## Complexity

Worst Case Time Complexity [ Big-O ]: **O(n log n)**  
Best Case Time Complexity [Big-omega]: **O(n log n)**  
Average Time Complexity [Big-theta]: **O(n log n)**  
Space Complexity: **O(n)**

## Algorithm

* Build a max heap from the input data.
* At this point, the largest item is stored at the root of the heap. Replace it with the last item of the heap followed by reducing the size of heap by 1. Finally, heapify the root of the tree.
* Repeat step 2 while size of heap is greater than 1.
