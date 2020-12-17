package main

import (
	"fmt"
	"math"

	"github.com/danrusei/algorithms_with_Go/graph"
)

type spf struct {
	*graph.Graph
}

// this code calculates shortest distance, but doesn’t calculate the path information.
func (g *spf) dijkstra(n *graph.Node) {

	nrNodes := len(g.Nodes)

	// dist[G] holds the shortest distance from source(in our case node n) to node G
	dist := make(map[*graph.Node]int)

	// initialize all distances as int Max and the source node with 0
	for _, node := range g.Nodes {
		dist[node] = math.MaxInt32
	}

	dist[n] = 0

	// visited[G] will be true if Node G is included in shortest path tree
	// or shortest distance from source to G is finalized
	visited := make(map[*graph.Node]bool)

	//find shortest path to all Nodes
	for i := 0; i < nrNodes-1; i++ {

		//Pick the minimum distance Node to the source from the set of vertices not yet processed.
		nodeX := g.minDistance(dist, visited)

		// Mark the picked vertex as processed
		visited[nodeX] = true

		//update dist value of the adjacent Nodes of the picked Node
		for _, nodeY := range g.Nodes {

			if visited[nodeY] || dist[nodeX] == math.MaxInt32 {
				continue
			}

			//if there is an link between nodeX to nodeY, update dist of adjacent nodes value
			for _, link := range g.Edges[nodeX] {
				if link.ToNode == nodeY && link.Cost+dist[nodeX] < dist[nodeY] {
					dist[nodeY] = dist[nodeX] + link.Cost
				}
			}
		}
	}

	fmt.Println(dist)

}

// A utility function to find the Node with minimum distance value, from
// the set of Nodes not yet included in shortest path tree
func (g *spf) minDistance(dist map[*graph.Node]int, visited map[*graph.Node]bool) *graph.Node {
	min := math.MaxInt32
	minNode := &graph.Node{}

	for _, node := range g.Nodes {
		if !visited[node] && (dist[node] <= min) {
			min = dist[node]
			minNode = node
		}
	}

	return minNode
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
