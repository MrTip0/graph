package graph

import (
	"math"

	"github.com/MrTip0/graph/queue"
)

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

func (g Graph) BFS(a, b string) int {
	border, bordern := queue.New(), queue.New()
	visited := make([]string, 0)

	d := 1

	border.Enqueue(a)

	for border.Len() > 0 {
		n := border.Dequeue()

		f := g.getByName(n.(string))

		for _, ele := range f.Nears {
			if ele.Dest == b {
				return d
			}
			if !contains(ele.Dest, visited) && !bordern.Contains(ele.Dest) {
				bordern.Enqueue(ele.Dest)
			}
		}

		visited = append(visited, f.Name)

		if border.Len() == 0 {
			d++
			border = bordern
			bordern = queue.New()
		}
	}
	return -1
}

func (g Graph) Dijkstra(a, b string) int {
	type data struct {
		steps int
		node  Node
	}

	vertexs := make([]data, len(g.V))
	l := len(g.V)

	strToPos := func(v string) int {
		r := -1
		for i := 0; i < l && r == -1; i++ {
			if vertexs[i].node.Name == v {
				r = i
			}
		}
		return r
	}

	border := queue.New()

	for i := 0; i < l; i++ {
		vertexs[i] = data{steps: math.MaxInt, node: g.V[i]}
	}

	if in := strToPos(a); in == -1 {
		return -2
	} else {
		vertexs[in].steps = 0
	}

	border.Enqueue(a)

	for border.Len() > 0 {
		el := strToPos(border.Dequeue().(string))

		for _, near := range vertexs[el].node.Nears {
			if pn := strToPos(near.Dest); vertexs[pn].steps > vertexs[el].steps+near.Peso {
				vertexs[pn].steps = vertexs[el].steps + near.Peso
				if !border.Contains(vertexs[pn].node.Name) {
					border.Enqueue(vertexs[pn].node.Name)
				}
			}
		}
	}

	end := strToPos(b)
	if end != -1 {
		return vertexs[end].steps
	}
	return -1
}

func contains(val string, l []string) bool {
	r := false

	for _, ele := range l {
		if ele == val {
			r = true
			break
		}
	}

	return r
}
