package graph

import (
	"reflect"
	"testing"
)

var g Graph

func TestGraph(t *testing.T) {
	createGraph()
	g.String()

	t.Run("Test AddNode()", func(t *testing.T) {
		nG := Node{"G"}
		g.AddNode(&nG)

		want := []string{"A", "B", "C", "D", "E", "F", "G"}
		got := []string{}

		for _, node := range g.Nodes {
			got = append(got, node.Value)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test AddEdge()", func(t *testing.T) {
		g.AddEdge(g.Nodes[5], g.Nodes[6])
		found := false

		for _, link := range g.Edges[g.Nodes[5]] {
			if link.ToNode == g.Nodes[6] {
				found = true
			}
		}
		if !found {
			t.Errorf("Edge creation has failed")
		}
	})

	t.Run("Test RemoveEdge()", func(t *testing.T) {
		g.RemoveEdge(g.Nodes[5], g.Nodes[6])
		found := false

		for _, link := range g.Edges[g.Nodes[5]] {
			if link.ToNode == g.Nodes[6] {
				found = true
			}
		}
		if found {
			t.Errorf("Edge deletion has failed")
		}
	})

	t.Run("Test RemoveNode()", func(t *testing.T) {
		g.RemoveNode(g.Nodes[6])

		want := []string{"A", "B", "C", "D", "E", "F"}
		got := []string{}

		for _, node := range g.Nodes {
			got = append(got, node.Value)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func createGraph() {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
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
}
