package graph

import (
	"testing"
)

func TestFindNode(t *testing.T) {
	g := createGraph()

	var tests = []struct {
		input  string
		output *Node
	}{
		{"A", g.Nodes[0]},
		{"D", g.Nodes[3]},
		{"invalid", nil},
	}

	for _, test := range tests {
		result := g.FindNode(test.input)
		if result != test.output {
			t.Errorf("FindNode: expected %v, got %v", test.output, result)
		}
	}
}

func TestBFS_FromA(t *testing.T) {
	g := createGraph()

	nodes := g.BFS(g.Nodes[0])
	expected := []string{"A", "B", "C", "D", "E", "F"}
	for i := range nodes {
		if nodes[i].Value != expected[i] {
			t.Errorf("BFS at index %d: expected %v, got %v", i, expected[i], nodes[i].Value)
		}
	}
}

func TestBFS_FromB(t *testing.T) {
	g := createGraph()

	nodes := g.BFS(g.Nodes[1])
	expected := []string{"B", "D"}

	for i := range nodes {
		if nodes[i].Value != expected[i] {
			t.Errorf("BFS at index %d: expected %v, got %v", i, expected[i], nodes[i].Value)
		}
	}
}

func TestBFS_FromC(t *testing.T) {
	g := createGraph()

	nodes := g.BFS(g.Nodes[2])
	expected := []string{"C", "E", "F"}
	for i := range nodes {
		if nodes[i].Value != expected[i] {
			t.Errorf("BFS at index %d: expected %v, got %v", i, expected[i], nodes[i].Value)
		}
	}
}

func createGraph() *Graph {
	g := New()
	a := Node{Value: "A"}
	b := Node{Value: "B"}
	c := Node{Value: "C"}
	d := Node{Value: "D"}
	e := Node{Value: "E"}
	f := Node{Value: "F"}

	a.Associate(&b, &c)
	b.Associate(&d)
	c.Associate(&e, &f)

	g.Add(&a, &b, &c, &d, &e, &f)

	return g
}
