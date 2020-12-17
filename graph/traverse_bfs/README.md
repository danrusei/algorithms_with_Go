# Breadth First Search Traversal

Source: [Wikipedia](https://en.wikipedia.org/wiki/Breadth-first_search)
Problem description: [GeeksforGeeks](https://www.geeksforgeeks.org/breadth-first-search-or-bfs-for-a-graph/)

Breadth-first search (BFS) is an algorithm for traversing or searching tree or graph data structures. It starts at the tree root (or some arbitrary node of a graph, sometimes referred to as a `search key`), and explores all of the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level.

It uses the opposite strategy of depth-first search, which instead explores the node branch as far as possible before being forced to backtrack and expand other nodes.

## Algorithm

This is a non-recursive implementation.

* it uses a queue (First In First Out) instead of a stack which contains the frontier along which the algorithm is currently searching.
* it checks whether a Node (vertex) has been discovered before enqueueing it
* Nodes are labelled as discovered by storing them in a set, or by an attribute on each node, depending on the implementation(in the current implementation we use the `map[*Node]bool`).

***Note that the word Node is usually interchangeable with the word Vertex.***

## Example

![bfs image](Animated_BFS.gif)
