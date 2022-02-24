package graphs

import "github.com/MrTip0/graph/containers"

func (g Graph) BFS(a, b string) int {
	edge, edgen := containers.NewQueue(), containers.NewQueue()
	visited := containers.NewSet()

	d := 1

	edge.Enqueue(a)

	for edge.Len() > 0 {
		n := edge.Dequeue()

		f := g.getByName(n.(string))

		for _, ele := range f.Nears {
			if ele.Dest == b {
				return d
			}
			if !visited.Exists(ele.Dest) && !edgen.Contains(ele.Dest) && !edge.Contains(ele.Dest) {
				edgen.Enqueue(ele.Dest)
			}
		}

		visited.Add(f.Name)

		if edge.Len() == 0 {
			d++
			edge = edgen
			edgen = containers.NewQueue()
		}
	}
	return -1
}
