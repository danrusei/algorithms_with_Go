package graph

import (
	"errors"
	"fmt"
)

// Graph represents the graph structure, which has the list of Nodes
// and the Edges(links) between the Nodes
type Graph struct {
	Nodes []*Node
	Edges map[*Node][]*Edge
}

// Node represents a Graph Node (vertex)
// where Value is a discretionary value associated to the node
type Node struct {
	Value string
}

// to be able to print the Node
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// Edge represents a link between two nodes (vertices)
// the Cost field is the weight between nodes, distance etc
// the ToNode field is the other end of the link (as the Edges map key is the first end)
type Edge struct {
	Cost   int
	ToNode *Node
}

// to be able to print the Edge
func (e *Edge) String() string {
	return fmt.Sprintf("%v", e.ToNode.Value)
}

// Adjancent : tests whether there is an edge from the node x to the node y;
func (g *Graph) Adjancent(x *Node, y *Node) bool {
	if _, ok := g.Edges[x]; ok {
		for _, link := range g.Edges[x] {
			if link.ToNode == y {
				return true
			}
		}
	}
	return false
}

// Neighbors : lists all nodes y such that there is an edge from the node x to the node y;
func (g *Graph) Neighbors(node *Node) ([]*Node, bool) {
	var resultedNodes []*Node

	if links, ok := g.Edges[node]; ok {
		for _, link := range links {
			resultedNodes = append(resultedNodes, link.ToNode)

		}
		return resultedNodes, true
	}
	return nil, false
}

// AddNode : adds a node if it's not there;
func (g *Graph) AddNode(node *Node) error {
	//check if the node is already there
	for _, foundNode := range g.Nodes {
		if foundNode == node {
			return errors.New("Node is already present in the list")
		}
	}
	g.Nodes = append(g.Nodes, node)
	return nil
}

// RemoveNode : removes the node by ID, if it is there;
func (g *Graph) RemoveNode(node *Node) bool {
	for idx, foundnode := range g.Nodes {
		if foundnode == node {
			g.Nodes = append(g.Nodes[:idx], g.Nodes[idx+1:]...)
			return true
		}
	}
	return false
}

// AddEdge : adds the edge from the node x to the node y, if it is not there;
func (g *Graph) AddEdge(x, y *Node) error {
	foundNode1, foundNode2 := false, false
	for _, node := range g.Nodes {
		if x == node {
			foundNode1 = true
		} else if y == node {
			foundNode2 = true
		}
	}
	if !foundNode1 || !foundNode2 {
		return errors.New("The nodes couldn't be found")
	}

	if g.Edges == nil {
		g.Edges = make(map[*Node][]*Edge)
	}

	g.Edges[x] = append(g.Edges[x], &Edge{ToNode: y})
	g.Edges[y] = append(g.Edges[y], &Edge{ToNode: x})

	return nil
}

// RemoveEdge : removes the edge from the node x to the node y, if it is there;
func (g *Graph) RemoveEdge(x, y *Node) error {
	for idx, link := range g.Edges[x] {
		if link.ToNode == y {
			g.Edges[x] = append(g.Edges[x][:idx], g.Edges[x][idx+1:]...)
		}
	}
	for idx, link := range g.Edges[y] {
		if link.ToNode == x {
			g.Edges[y] = append(g.Edges[y][:idx], g.Edges[y][idx+1:]...)
		}
	}
	return nil
}

// GetNodeValue : returns the value associated with the node x;
func (g *Graph) GetNodeValue(x *Node) string {
	return x.Value
}

// SetNodeValue : sets the value associated with the node x to v.
func (g *Graph) SetNodeValue(x *Node, value string) {
	x.Value = value
}

// GetEdgeValue : returns the value associated with the edge (x, y);
func (g *Graph) GetEdgeValue(x *Node, y *Node) int {
	for _, link := range g.Edges[x] {
		if link.ToNode == y {
			return link.Cost
		}
	}
	return 0
}

// SetEdgeValue : sets the value associated with the edge (x, y) to v.
func (g *Graph) SetEdgeValue(x, y *Node, value int) bool {
	for _, link := range g.Edges[x] {
		if link.ToNode == y {
			link.Cost = value
			return true
		}
	}
	return false
}

// to be able to print the graph
func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + " -> "
		near := g.Edges[g.Nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}
