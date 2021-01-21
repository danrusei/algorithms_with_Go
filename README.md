# GO - Data Structures and Algorithms

Inspired by the **[Geeksforgeeks - Top 10 Algorithms in Interview Questions](https://www.geeksforgeeks.org/top-10-algorithms-in-interview-questions/amp/)** article, the intent of this repository is to solve these questions using the **Go Language**. GO is a great language choice for technical interviews and hopefully you can find these solutions to the common algorithms/problems easy to understand. 

Although this is an introductory material to algorithms and data structures, it assumes that you are familiar with GO programming language syntax and basic concepts. 

WIP, the descriptions of the below `unsolved yet` problems can be found in the [orginal article](https://www.geeksforgeeks.org/top-10-algorithms-in-interview-questions/amp/).  

> ***Contributions are welcomed - solve a problem and submit a PR***.  
> ***Contribution guidelines***
> * keep the consistency and document the code
> * optimize for readability and simplicity (not over-engineered, best performance is not in scope)

## [Graph](https://github.com/danrusei/algorithms_with_Go/tree/main/graph)

- [x] [Breadth First Search (BFS)](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/traverse_bfs)
- [x] [Depth First Search (DFS)](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/traverse_dfs)
- [x] [Shortest Path from source to all vertices (Dijkstra)](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/dijkstra)
- [] Shortest Path from every vertex to every other vertex (Floyd Warshall)
- [] To detect cycle in a Graph (Union Find)
- [] Minimum Spanning tree (Prim)
- [] Minimum Spanning tree (Kruskal)
- [] Topological Sort
- [x] [Boggle (Find all possible words in a board of characters)](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/boggle)
- [x] [Bridges in a Graph](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/bridges)

## [Linked List](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist)

- [x] [Insertion of a node in Linked List (On the basis of some constraints)](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/insert_node)
- [] Delete a given node in Linked List (under given constraints)
- [x] [Compare two strings represented as linked lists](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/compare_strings)
- [] Add Two Numbers Represented By Linked Lists
- [] Merge A Linked List Into Another Linked List At Alternate Positions
- [x] [Reverse A List In Groups Of Given Size](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/reverse_by_groups)
- [] Union And Intersection Of 2 Linked Lists
- [x] [Detect And Remove Loop In A Linked List](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/remove_loop)
- [] Merge Sort For Linked Lists
- [x] [Select A Random Node from A Singly Linked List](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/random_node)

## [Dynamic Programming](https://github.com/danrusei/algorithms_with_Go/tree/main/dynamic)

- [x] [Longest Common Subsequence](https://github.com/danrusei/algorithms_with_Go/tree/main/dynamic/longest_common_subsequence)
- [] Longest Increasing Subsequence
- [] Edit Distance
- [] Minimum Partition
- [x] [Ways to Cover a Distance](https://github.com/danrusei/algorithms_with_Go/tree/main/dynamic/cover_distance)
- [] Longest Path In Matrix
- [x] [Subset Sum Problem](https://github.com/danrusei/algorithms_with_Go/tree/main/dynamic/subset_sum)
- [] Optimal Strategy for a Game
- [] 0-1 Knapsack Problem
- [] Boolean Parenthesization Problem

## [Sorting And Searching](https://github.com/danrusei/algorithms_with_Go/tree/main/sorting)

- [] Binary Search
- [] Search an element in a sorted and rotated array
- [x] [Bubble Sort](https://github.com/danrusei/algorithms_with_Go/tree/main/sorting/bubble_sort)
- [x] [Insertion Sort](https://github.com/danrusei/algorithms_with_Go/tree/main/sorting/insertion_sort)
- [x] [Merge Sort](https://github.com/danrusei/algorithms_with_Go/tree/main/sorting/merge_sort)
- [x] [Heap Sort (Binary Heap)](https://github.com/danrusei/algorithms_with_Go/tree/main/sorting/heap_sort)
- [x] [Quick Sort](https://github.com/danrusei/algorithms_with_Go/tree/main/sorting/quick_sort)
- [] Interpolation Search
- [] Find Kth Smallest/Largest Element In Unsorted Array
- [] Given a sorted array and a number x, find the pair in array whose sum is closest to x

## [Tree / Binary Search Tree](https://github.com/danrusei/algorithms_with_Go/tree/main/binary_tree)

- [x] [Find Minimum Depth of a Binary Tree](https://github.com/danrusei/algorithms_with_Go/tree/main/binary_tree/minimum_depth)
- [] Maximum Path Sum in a Binary Tree
- [x] [Check if a given array can represent Preorder Traversal of Binary Search Tree](https://github.com/danrusei/algorithms_with_Go/tree/main/binary_tree/preorder_traversal)
- [] Check whether a binary tree is a full binary tree or not
- [x] [Bottom View Binary Tree](https://github.com/danrusei/algorithms_with_Go/tree/main/binary_tree/bottom_view)
- [] Print Nodes in Top View of Binary Tree
- [x] [Remove nodes on root to leaf paths of length < K](https://github.com/danrusei/algorithms_with_Go/tree/main/binary_tree/remove_nodes)
- [] Lowest Common Ancestor in a Binary Search Tree
- [] Check if a binary tree is subtree of another binary tree
- [x] [Reverse alternate levels of a perfect binary tree](https://github.com/danrusei/algorithms_with_Go/tree/main/binary_tree/reverse_alternate)

## [Number Theory](https://github.com/danrusei/algorithms_with_Go/tree/main/numbers)

- [] Modular Exponentiation
- [x] [Modular multiplicative inverse](https://github.com/danrusei/algorithms_with_Go/tree/main/numbers/multiplicative)
- [] Primality Test | Set 2 (Fermat Method)
- [] Euler’s Totient Function
- [x] [Sieve of Eratosthenes](https://github.com/danrusei/algorithms_with_Go/tree/main/numbers/eratosthenes)
- [] Convex Hull
- [] Basic and Extended Euclidean algorithms
- [x] [Segmented Sieve](https://github.com/danrusei/algorithms_with_Go/tree/main/numbers/segmented)
- [] Chinese remainder theorem
- [] Lucas Theorem

## [BIT Manipulation](https://github.com/danrusei/algorithms_with_Go/tree/main/bitwise)

- [x] [Maximum Subarray XOR](https://github.com/danrusei/algorithms_with_Go/tree/main/bitwise/max_xor)
- [] Magic Number
- [] Sum of bit differences among all pairs
- [x] [Swap All Odds And Even Bits](https://github.com/danrusei/algorithms_with_Go/tree/main/bitwise/swapp_odd_even)
- [] Find the element that appears once
- [x] [Binary representation of a given number](https://github.com/danrusei/algorithms_with_Go/tree/main/bitwise/decimal_to_binary)
- [] Count total set bits in all numbers from 1 to n
- [x] [Rotate bits of a number](https://github.com/danrusei/algorithms_with_Go/tree/main/bitwise/rotate_bits)
- [] Count number of bits to be flipped to convert A to B
- [x] [Find Next Sparse Number](https://github.com/danrusei/algorithms_with_Go/tree/main/bitwise/next_sparse)

## [String / Array](https://github.com/danrusei/algorithms_with_Go/tree/main/strings)

- [x] [Reverse an array without affecting special characters](https://github.com/danrusei/algorithms_with_Go/tree/main/strings/reverse_alpha)
- [] All Possible Palindromic Partitions
- [] Count triplets with sum smaller than a given value
- [] Convert array into Zig-Zag fashion
- [] Generate all possible sorted arrays from alternate elements of two given sorted arrays
- [] Pythagorean Triplet in an array
- [x] [Length of the largest subarray with contiguous elements](https://github.com/danrusei/algorithms_with_Go/tree/main/strings/largest_subarray)
- [] Find the smallest positive integer value that cannot be represented as sum of any subset of a given array
- [x] [Smallest subarray with sum greater than a given value](https://github.com/danrusei/algorithms_with_Go/tree/main/strings/smallest_subarray)
- [] Stock Buy Sell to Maximize Profit

## [Other Data Structures](https://github.com/danrusei/algorithms_with_Go/tree/main/other_ds)

- [x] [Stack](https://github.com/danrusei/algorithms_with_Go/tree/main/other_ds/stack)
- [x] [Queue](https://github.com/danrusei/algorithms_with_Go/tree/main/other_ds/queue)
- [] HashMap

