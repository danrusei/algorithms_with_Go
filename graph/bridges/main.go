package main

import (
	"fmt"

	"github.com/danrusei/algorithms_with_Go/graph"
)

type bridges struct {
	*graph.Graph
}

func (g *bridges) findBridge() {

	totalNodes := len(g.Nodes)

	for _, node := range g.Nodes {
		for _, edge := range g.Edges[node] {
			//remove the (u, v) edge  from graph
			g.RemoveEdge(node, edge.ToNode)

			//check if bfs traversal returns the same number of nodes as expected
			if g.BFStraverse(node) != totalNodes {
				fmt.Printf("Edge %v --- %v is a bridge\n", edge.ToNode, node)
			}

			// add again the (u, v) edge to graph, in order to check the other edges as well
			g.AddEdge(node, edge.ToNode)
		}
	}
}

// BFS(Breadth First Search) traversal of the Graph
// check out https://github.com/danrusei/algorithms_with_Go/tree/main/graph/traverse_bfs
//for detailed explanation of the algorithm
func (g *bridges) BFStraverse(n *graph.Node) int {
	// Create a queue for BFS
	queue := make([]*graph.Node, 0, len(g.Nodes))

	queue = append(queue, n)
	visited := make(map[*graph.Node]bool)

	for len(queue) > 0 {
		node := queue[0]
		//pop out the Node from the queue
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
	return len(visited)
}

func main() {
	// create nodes
	n0 := graph.Node{Value: "0"}
	n1 := graph.Node{Value: "1"}
	n2 := graph.Node{Value: "2"}
	n3 := graph.Node{Value: "3"}
	n4 := graph.Node{Value: "4"}

	// create a new graph instance
	g := new(graph.Graph)

	// add nodes to graph
	g.AddNode(&n0)
	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)

	// add the edges(links) between nodes
	g.AddEdge(&n0, &n1)
	g.AddEdge(&n0, &n3)
	g.AddEdge(&n0, &n2)
	g.AddEdge(&n1, &n2)
	g.AddEdge(&n3, &n4)

	s := &bridges{g}
	s.findBridge()
}
