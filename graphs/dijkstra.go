package graphs

import (
	"math"

	"github.com/MrTip0/graph/queue"
)

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
