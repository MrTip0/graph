package graphs

import "github.com/MrTip0/graph/queue"

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
