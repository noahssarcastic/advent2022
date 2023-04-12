package main

type graph struct {
	nodes []node
}

func newGraph() *graph {
	return &graph{make([]node, 0)}
}

func (g *graph) v() int {
	return len(g.nodes)
}

func (g *graph) get(i int) node {
	return g.nodes[i]
}

func (g *graph) addNode(pressure int, label string, adj []int) {
	g.nodes = append(g.nodes, node{
		pressure: pressure,
		label:    label,
		adjacent: adj,
	})
}

type node struct {
	pressure int
	label    string
	adjacent []int
}
