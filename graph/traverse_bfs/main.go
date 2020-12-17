package main

import (
	"fmt"

	"github.com/danrusei/algorithms_with_Go/graph"
)

type bfs struct {
	*graph.Graph
}

// BFS(Breadth First Search) traversal of the Graph
func (g *bfs) BFStraverse() {
	// Create a queue for BFS
	queue := make([]*graph.Node, 0, len(g.Nodes))
	//select a starting Node, in our case first node defined
	n := g.Nodes[0]

	queue = append(queue, n)
	visited := make(map[*graph.Node]bool)

	for len(queue) > 0 {
		node := queue[0]
		//print and the pop out the Node from the queue
		fmt.Println(node)
		queue = queue[1:]
		// add the Node to visited map, ensuring that it will not be added again to the queue
		// this prevents to loop endlesssly through the nodes
		visited[node] = true

		//edges are the links to other Nodes of the searching Node
		for _, edge := range g.Edges[node] {
			if !visited[edge.ToNode] {
				newNode := edge.ToNode
				queue = append(queue, newNode)
				visited[newNode] = true
			}

		}

	}
}

func main() {
	nA := graph.Node{Value: "A"}
	nB := graph.Node{Value: "B"}
	nC := graph.Node{Value: "C"}
	nD := graph.Node{Value: "D"}
	nE := graph.Node{Value: "E"}
	nF := graph.Node{Value: "F"}

	g := new(graph.Graph)

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nA)
	g.AddEdge(&nD, &nA)

	s := &bfs{g}
	s.BFStraverse()
}
