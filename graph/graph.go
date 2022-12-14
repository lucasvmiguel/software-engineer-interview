package graph

type Node struct {
	Value string
	Edges []*Node
}

type Graph struct {
	Nodes []*Node
}

func New() *Graph {
	return &Graph{}
}

func (n *Node) Associate(nodes ...*Node) {
	n.Edges = append(n.Edges, nodes...)

	// for _, node := range nodes {
	// 	node.Edges = append(node.Edges, n)
	// }
}

func (g *Graph) Add(nodes ...*Node) {
	g.Nodes = append(g.Nodes, nodes...)
}

func (g *Graph) FindNode(value string) *Node {
	for _, node := range g.Nodes {
		if node.Value == value {
			return node
		}
	}

	return nil
}

func (g *Graph) BFS(node *Node) []*Node {
	queue := []*Node{node}
	visited := []*Node{}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if !contains(visited, n) {
			visited = append(visited, n)
			queue = append(queue, n.Edges...)
		}
	}

	return visited
}

func contains(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n.Value == node.Value {
			return true
		}
	}

	return false
}
