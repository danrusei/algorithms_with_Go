package graph

// Graph represents the graph structure
type Graph struct {
	Nodes []*Node
}

// Node represents a Graph Node (vertex)
// where id is the unique id of the node,
// value is a discretionary value associated to the node
// and edges are the node links to the other nodes
type Node struct {
	ID    int
	Value int
	Edges map[int]Edge
}

// Edge represents a link between two nodes (vertices)
// with the cost that can represent a weight between nodes, distance etc
type Edge struct {
	Cost int
}

// Adjancent : tests whether there is an edge from the node x to the node y;
func (g *Graph) Adjancent(x *Node, y *Node) bool {
	return false
}

// Neighbors : lists all nodes y such that there is an edge from the node x to the node y;
func (g *Graph) Neighbors(x *Node) []*Node {
	return nil
}

// AddNode : adds the node x, if it is not there;
func (g *Graph) AddNode(x *Node) bool {
	return false
}

// RemoveNode : removes the node x, if it is there;
func (g *Graph) RemoveNode(x *Node) bool {
	return false
}

// AddEdge : adds the edge from the node x to the node y, if it is not there;
func (g *Graph) AddEdge(x *Node, y *Node) bool {
	return false
}

// RemoveEdge : removes the edge from the node x to the node y, if it is there;
func (g *Graph) RemoveEdge(x *Node, y *Node) bool {
	return false
}

// GetNodeValue : returns the value associated with the node x;
func (g *Graph) GetNodeValue(x *Node) int {
	return 0
}

// SetNodeValue : sets the value associated with the node x to v.
func (g *Graph) SetNodeValue(x *Node, value int) bool {
	return false
}

// GetEdgeValue : returns the value associated with the edge (x, y);
func (g *Graph) GetEdgeValue(x *Node, y *Node) int {
	return 0
}

// SetEdgeValue : sets the value associated with the edge (x, y) to v.
func (g *Graph) SetEdgeValue(x *Node, y *Node, value int) bool {
	return false
}
