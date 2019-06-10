package main

// Balance calc balance
func (g Graph) Balance() bool {
	for x := 0; x < len(g.nodes)-2; x++ {
		for y := x + 1; y < len(g.nodes)-1; y++ {
			for z := y + 1; z < len(g.nodes); z++ {
				if !g.BalanceTriplet(g.nodes[x], g.nodes[y], g.nodes[z]) {
					return false
				}
			}
		}
	}
	return true
}

// BalanceTriplet check if local balanced
func (g Graph) BalanceTriplet(x *Node, y *Node, z *Node) bool {
	if g.CheckTripleConnection(x, y, z) {
		score := g.CheckConnection(x, y) +
			g.CheckConnection(x, z) +
			g.CheckConnection(y, z)
		if score != 3 && score != -1 {
			return false
		}
	}
	return true
}

// CheckConnection between nodes
func (g Graph) CheckConnection(x *Node, y *Node) int {
	for _, n := range g.friend[*x] {
		if n == y {
			return 1
		}
	}
	for _, n := range g.enemy[*x] {
		if n == y {
			return -1
		}
	}
	return 0
}

// CheckTripleConnection between three nodes
func (g Graph) CheckTripleConnection(x *Node, y *Node, z *Node) bool {
	return g.CheckConnection(x, y) != 0 &&
		g.CheckConnection(x, z) != 0 &&
		g.CheckConnection(y, z) != 0
}
