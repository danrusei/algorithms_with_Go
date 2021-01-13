# GO - Data Structures and Algorithms

Inspired by the **[Geeksforgeeks - Top 10 Algorithms in Interview Questions](https://www.geeksforgeeks.org/top-10-algorithms-in-interview-questions/amp/)** article, the intent of this repository is to solve these questions using the **Go Language**.  GO is a great language choice for technical interviews and hopefully you can find the solutions to the common algorithms/problems easy to understand. 

Although this is an introductory course to algorithms and data structures, it assumes that you are familiar with GO programming language syntax and basic concepts. 

WIP, the descriptions of the below `unsolved yet` problems can be found in the [orginal article](https://www.geeksforgeeks.org/top-10-algorithms-in-interview-questions/amp/).  

> ***Contributions are welcomed - submit a PR***.  
> ***Contribution guidelines***
> * keep the consistency and document the code
> * keep it simple, easy to read and understand (not over-engineered, best performance is not in scope)

### [Queue](https://github.com/danrusei/algorithms_with_Go/tree/main/queue)

### [Stack](https://github.com/danrusei/algorithms_with_Go/tree/main/stack)

## [Graph](https://github.com/danrusei/algorithms_with_Go/tree/main/graph)

- [x] [Breadth First Search (BFS)](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/traverse_bfs)
- [x] [Depth First Search (DFS)](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/traverse_dfs)
- [x] [Shortest Path from source to all vertices (Dijkstra)](https://github.com/danrusei/algorithms_with_Go/tree/main/graph/dijkstra)
- [] Shortest Path from every vertex to every other vertex (Floyd Warshall)
- [] To detect cycle in a Graph (Union Find)
- [] Minimum Spanning tree (Prim)
- [] Minimum Spanning tree (Kruskal)
- [] Topological Sort
- [] Boggle (Find all possible words in a board of characters)
- [] Bridges in a Graph

## [Linked List](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist)

- [x] [Insertion of a node in Linked List (On the basis of some constraints)](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/insert_node)
- [] Delete a given node in Linked List (under given constraints)
- [] Compare two strings represented as linked lists
- [] Add Two Numbers Represented By Linked Lists
- [] Merge A Linked List Into Another Linked List At Alternate Positions
- [x] [Reverse A List In Groups Of Given Size](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/reverse_by_groups)
- [] Union And Intersection Of 2 Linked Lists
- [x] [Detect And Remove Loop In A Linked List](https://github.com/danrusei/algorithms_with_Go/tree/main/linkedlist/remove_loop)
- [] Merge Sort For Linked Lists
- [] Select A Random Node from A Singly Linked List

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
- [] Remove nodes on root to leaf paths of length < K
- [] Lowest Common Ancestor in a Binary Search Tree
- [] Check if a binary tree is subtree of another binary tree
- [] Reverse alternate levels of a perfect binary tree

## [Number Theory](https://github.com/danrusei/algorithms_with_Go/tree/main/numbers)

- [] Modular Exponentiation
- [] Modular multiplicative inverse
- [] Primality Test | Set 2 (Fermat Method)
- [] Euler’s Totient Function
- [] Sieve of Eratosthenes
- [] Convex Hull
- [] Basic and Extended Euclidean algorithms
- [] Segmented Sieve
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
- [] Rotate bits of a number
- [] Count number of bits to be flipped to convert A to B
- [] Find Next Sparse Number

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

