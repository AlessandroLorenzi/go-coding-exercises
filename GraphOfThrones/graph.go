package main

// Graph is a struct
type Graph struct {
	nodes  []*Node
	friend map[Node][]*Node
	enemy  map[Node][]*Node
}

// Node is a struct
type Node struct {
	value string
}

// NewNode Add a node to graph
func (g *Graph) NewNode(name string) *Node {
	if g.NodeExists(name) {
		return g.GetNode(name)
	}
	n := Node{name}

	g.nodes = append(g.nodes, &n)
	g.friend[n] = make([]*Node, 0)
	g.enemy[n] = make([]*Node, 0)

	return &n
}

// NodeExists return true if node exists, false if not
func (g *Graph) NodeExists(name string) bool {
	for _, node := range g.nodes {
		if (*node).value == name {
			return true
		}
	}
	return false
}

// GetNode return true if node exists, false if not
func (g *Graph) GetNode(name string) *Node {
	for _, node := range g.nodes {
		if (*node).value == name {
			return node
		}
	}
	return nil
}

// AddEdge between people
func (g *Graph) AddEdge(fn *Node, sn *Node, friends bool) {
	if friends {
		g.friend[*fn] = append(g.friend[*fn], sn)
		g.friend[*sn] = append(g.friend[*sn], fn)
	} else {
		g.enemy[*fn] = append(g.enemy[*fn], sn)
		g.enemy[*sn] = append(g.enemy[*sn], fn)
	}
}

// String export graph as a string
func (g *Graph) String() string {
	ret := "Nodes: "
	for _, n := range g.nodes {
		ret = ret + "\n   " + (*n).value
	}
	ret = ret + "\n\n"
	for k, v := range g.friend {
		for _, i := range v {
			ret = ret + k.value + " ++ " + (*i).value + "\n"
		}
	}
	for k, v := range g.enemy {
		for _, i := range v {
			ret = ret + k.value + " -- " + (*i).value + "\n"
		}
	}
	return ret
}
