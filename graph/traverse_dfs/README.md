# Depth First Search Traversal

Source: [Wikipedia](https://en.wikipedia.org/wiki/Breadth-first_search)
Problem description: [GeeksforGeeks](https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/)

Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph) and explores as far as possible along each branch before backtracking.

## Algorithm

* create a recursive function that takes the index of node and a visited array.
* mark the current node as visited and print the node.
* traverse all the adjacent and unmarked nodes and call the recursive function with index of adjacent node.

## Example

![depth first search](Depth-First-Search.gif)
