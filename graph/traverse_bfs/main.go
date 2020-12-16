package main

import (
	"github.com/danrusei/algorithms_with_Go/graph"
	"github.com/danrusei/algorithms_with_Go/queue"
)

type bfs struct {
	*graph.Graph
}

func (g *bfs) BFStraverse() {
	s := new(queue.QueueString)

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
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)

	s := &bfs{g}
	s.BFStraverse()

}
