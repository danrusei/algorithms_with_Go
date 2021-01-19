package main

import (
	"fmt"

	"github.com/danrusei/algorithms_with_Go/graph"
)

type boggle struct {
	*graph.Graph
}

var words = []string{"GEEKS", "FOR", "QUIZ", "GO"}

func (g *boggle) isWord(word string) bool {

	//find the node value equal with the first letter of the word
	first := []*graph.Node{}
	for _, node := range g.Nodes {
		if node.Value == string(word[0]) {
			first = append(first, node)
		}
	}

	// a modified breadth first search traversal of the graph
	// as we note as visited only the Nodes with the values that satisfy the condition:
	// Node.Value == string(word[wIndex])
	visited := make(map[*graph.Node]bool)
	for _, node := range first {
		visited[node] = true
		wIndex := 0
		nextNode := node

		for {
			if wIndex == len(word)-1 {
				return true
			}
			wIndex++
			for _, edge := range g.Edges[nextNode] {
				if !visited[edge.ToNode] {
					if edge.ToNode.Value == string(word[wIndex]) {
						nextNode = edge.ToNode
						break
					}
				}
			}
			// it means that the nextNode has not changed, therefore not able to find another
			// suitable Node in the above iteration
			if visited[nextNode] {
				return false
			}
			visited[nextNode] = true
		}
	}
	return false
}

func main() {
	// create nodes
	nG := graph.Node{Value: "G"}
	nI := graph.Node{Value: "I"}
	nZ := graph.Node{Value: "Z"}
	nU := graph.Node{Value: "U"}
	nE := graph.Node{Value: "E"}
	nK := graph.Node{Value: "K"}
	nQ := graph.Node{Value: "Q"}
	nS := graph.Node{Value: "S"}
	nE2 := graph.Node{Value: "E"}

	// create a new graph instance
	g := new(graph.Graph)

	// add nodes to graph
	g.AddNode(&nG)
	g.AddNode(&nI)
	g.AddNode(&nZ)
	g.AddNode(&nU)
	g.AddNode(&nE)
	g.AddNode(&nK)
	g.AddNode(&nQ)
	g.AddNode(&nS)
	g.AddNode(&nE2)

	// add the edges(links) between nodes
	g.AddEdge(&nG, &nI)
	g.AddEdge(&nG, &nE)
	g.AddEdge(&nG, &nU)
	g.AddEdge(&nI, &nG)
	g.AddEdge(&nI, &nU)
	g.AddEdge(&nI, &nE)
	g.AddEdge(&nI, &nK)
	g.AddEdge(&nI, &nZ)
	g.AddEdge(&nZ, &nI)
	g.AddEdge(&nZ, &nE)
	g.AddEdge(&nZ, &nK)
	g.AddEdge(&nU, &nG)
	g.AddEdge(&nU, &nI)
	g.AddEdge(&nU, &nE)
	g.AddEdge(&nU, &nS)
	g.AddEdge(&nU, &nQ)
	g.AddEdge(&nE, &nG)
	g.AddEdge(&nE, &nI)
	g.AddEdge(&nE, &nZ)
	g.AddEdge(&nE, &nU)
	g.AddEdge(&nE, &nK)
	g.AddEdge(&nE, &nQ)
	g.AddEdge(&nE, &nS)
	g.AddEdge(&nE, &nE2)
	g.AddEdge(&nK, &nZ)
	g.AddEdge(&nK, &nI)
	g.AddEdge(&nK, &nE)
	g.AddEdge(&nK, &nS)
	g.AddEdge(&nK, &nE2)
	g.AddEdge(&nQ, &nU)
	g.AddEdge(&nQ, &nE)
	g.AddEdge(&nQ, &nS)
	g.AddEdge(&nS, &nQ)
	g.AddEdge(&nS, &nU)
	g.AddEdge(&nS, &nE)
	g.AddEdge(&nS, &nK)
	g.AddEdge(&nS, &nE2)
	g.AddEdge(&nE2, &nS)
	g.AddEdge(&nE2, &nE)
	g.AddEdge(&nE2, &nK)

	s := &boggle{g}

	for _, word := range words {
		if s.isWord(word) {
			fmt.Printf("This word was found: %s\n", word)
		}
	}
}
