package main

import (
	"fmt"

	"github.com/danrusei/algorithms_with_Go/graph"
)

type dfs struct {
	*graph.Graph
}

// DFS(Depth First Search) traversal of the Graph
func (g *dfs) DFStraverse(node *graph.Node) {

	// visited map prevents endlesssly loops through the nodes
	visited := make(map[*graph.Node]bool)

	// Call the recursive helper function to print DFS traversal
	g.DFSUtil(node, visited)
}

func (g *dfs) DFSUtil(node *graph.Node, visited map[*graph.Node]bool) {

	// add the Node to visited map, ensuring that it will not looked again
	visited[node] = true
	fmt.Println(node)

	//edges are the current Node links to the other Nodes
	for _, edge := range g.Edges[node] {
		if !visited[edge.ToNode] {
			newNode := edge.ToNode
			g.DFSUtil(newNode, visited)
		}

	}

}

func main() {
	nA := graph.Node{Value: "A"}
	nB := graph.Node{Value: "B"}
	nC := graph.Node{Value: "C"}
	nD := graph.Node{Value: "D"}

	g := new(graph.Graph)

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nC)
	g.AddEdge(&nC, &nA)
	g.AddEdge(&nB, &nC)
	g.AddEdge(&nC, &nC)

	s := &dfs{g}
	s.DFStraverse(&nA)
}
