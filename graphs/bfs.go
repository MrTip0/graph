package graphs

import "github.com/MrTip0/graph/containers"

func (g Graph) BFS(a, b string) int {
	edge, edgen := queue.New(), queue.New()
	visited := make([]string, 0)

	d := 1

	edge.Enqueue(a)

	for edge.Len() > 0 {
		n := edge.Dequeue()

		f := g.getByName(n.(string))

		for _, ele := range f.Nears {
			if ele.Dest == b {
				return d
			}
			if !contains(ele.Dest, visited) && !edgen.Contains(ele.Dest) {
				edgen.Enqueue(ele.Dest)
			}
		}

		visited = append(visited, f.Name)

		if edge.Len() == 0 {
			d++
			edge = edgen
			edgen = queue.New()
		}
	}
	return -1
}
