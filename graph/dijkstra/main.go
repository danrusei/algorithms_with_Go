package main

import (
	"github.com/danrusei/algorithms_with_Go/graph"
)

type spf struct {
	*graph.Graph
}

func (g *spf) dijkstra(n *graph.Node) {

}

func main() {
	// create nodes
	nA := graph.Node{Value: "A"}
	nB := graph.Node{Value: "B"}
	nC := graph.Node{Value: "C"}
	nD := graph.Node{Value: "D"}
	nE := graph.Node{Value: "E"}
	nF := graph.Node{Value: "F"}
	nG := graph.Node{Value: "G"}
	nH := graph.Node{Value: "H"}
	nI := graph.Node{Value: "I"}

	// create a new graph instance
	g := new(graph.Graph)

	// add nodes to graph
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddNode(&nG)
	g.AddNode(&nH)
	g.AddNode(&nI)

	// add the edges(links) between nodes
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nC)
	g.AddEdge(&nB, &nD)
	g.AddEdge(&nD, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nC, &nF)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nG)
	g.AddEdge(&nG, &nH)
	g.AddEdge(&nD, &nH)
	g.AddEdge(&nF, &nH)
	g.AddEdge(&nG, &nI)
	g.AddEdge(&nH, &nI)

	// set the Cost of the Edges (Nodes)
	g.SetEdgeValue(&nA, &nB, 4)
	g.SetEdgeValue(&nA, &nC, 8)
	g.SetEdgeValue(&nB, &nC, 11)
	g.SetEdgeValue(&nB, &nD, 8)
	g.SetEdgeValue(&nD, &nE, 2)
	g.SetEdgeValue(&nC, &nE, 7)
	g.SetEdgeValue(&nC, &nF, 1)
	g.SetEdgeValue(&nE, &nF, 6)
	g.SetEdgeValue(&nD, &nG, 7)
	g.SetEdgeValue(&nG, &nH, 14)
	g.SetEdgeValue(&nD, &nH, 4)
	g.SetEdgeValue(&nF, &nH, 2)
	g.SetEdgeValue(&nG, &nI, 9)
	g.SetEdgeValue(&nH, &nI, 10)

	// run Dijkstra algorithm
	s := &spf{g}
	s.dijkstra(&nA)
}
