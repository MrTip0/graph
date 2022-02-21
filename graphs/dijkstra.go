package graphs

import (
	"math"

	"github.com/MrTip0/graph/containers"
)

func (g Graph) Dijkstra(a, b string) (int, string) {
	const empty_set = "\u00f8"

	type data struct {
		steps int
		node  Node
		path  string
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

	edge := containers.NewPriorityQueue()

	for i := 0; i < l; i++ {
		vertexs[i] = data{steps: math.MaxInt, node: g.V[i], path: g.V[i].Name}
	}

	if in := strToPos(a); in == -1 {
		return -2, empty_set
	} else {
		vertexs[in].steps = 0
	}

	edge.Enqueue(a, 0)

	for edge.Len() > 0 {
		el := strToPos(edge.Dequeue().(string))

		for _, near := range vertexs[el].node.Nears {
			if pn := strToPos(near.Dest); vertexs[pn].steps > vertexs[el].steps+near.Peso {
				vertexs[pn].steps = vertexs[el].steps + near.Peso
				vertexs[pn].path = vertexs[el].path + " -> " + vertexs[pn].node.Name
				if !edge.Contains(vertexs[pn].node.Name) {
					edge.Enqueue(vertexs[pn].node.Name, vertexs[pn].steps)
				} else {
					edge.UpdatePriority(vertexs[pn].node.Name, vertexs[pn].steps)
				}
			}
		}
	}

	end := strToPos(b)
	if st := vertexs[end].steps; end != -1 && st != math.MaxInt {
		return st, vertexs[end].path
	}
	return -1, empty_set
}
