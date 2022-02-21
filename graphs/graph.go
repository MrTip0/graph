package graphs

type Node struct {
	Name  string
	Nears []Line
}

type Line struct {
	Dest string
	Peso int
}

type Graph struct {
	V []Node
}

func New() Graph {
	return Graph{V: make([]Node, 0)}
}

func (n *Node) AddLine(dest string, peso int) {
	n.Nears = append(n.Nears, Line{Dest: dest, Peso: peso})
}

func (g *Graph) AddNode(name string) {
	l := len(g.V)

	// Check if the list already contains the node with that name
	for i := 0; i < l; i++ {
		if g.V[i].Name == name {
			return
		}
	}
	g.V = append(g.V, Node{Name: name, Nears: make([]Line, 0)})
}

func (g *Graph) AddLine(from, to string, peso int) {
	a, b := g.getByName(from), g.getByName(to)
	if a == nil || b == nil {
		return
	}
	a.Nears = append(a.Nears, Line{Dest: to, Peso: peso})
}

func (g *Graph) getByName(n string) *Node {
	l := len(g.V)
	for i := 0; i < l; i++ {
		if g.V[i].Name == n {
			return &g.V[i]
		}
	}
	return nil
}
