package graphs

import (
	"fmt"

	"github.com/MrTip0/graph/containers"
)

func (g Graph) DFS(begin string) string {
	visited := containers.NewSet()
	edge := containers.NewStack()

	edge.Push(begin)

EdgeLoop:
	for edge.Len() > 0 {
		n := edge.First()

		f := g.getByName(n.(string))

		for _, ele := range f.Nears {
			if !visited.Exists(ele.Dest) && !edge.Contains(ele.Dest) {
				edge.Push(ele.Dest)

				// If a node is putted in the edge this continue from it
				continue EdgeLoop
			}
		}

		edge.Pop()
		visited.Add(n.(string))
	}

	r := ""
	l := len(g.V)
	for i := 0; i < l && visited.Size() < l; i++ {
		if act := g.V[i].Name; !visited.Exists(act) {
			r = fmt.Sprintf("%s %s", r, act)
			visited.Add(act)
		}
	}
	return fmt.Sprintf("[%s ]", r)
}
