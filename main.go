package main

import (
	"fmt"
	"strconv"

	"github.com/MrTip0/graph/graphs"
)

func main() {
	comm := 1
	g := graphs.New()

	for comm != 0 {
		fmt.Println("0 - Exit")
		fmt.Println("1 - Add Node")
		fmt.Println("2 - Add Line")
		fmt.Println("3 - BFS")
		fmt.Println("4 - Dijkstra")
		fmt.Println()

		var com string
		var err error

		fmt.Print("> ")
		fmt.Scanf("%s", &com)

		comm , err = strconv.Atoi(com)
		if err != nil {
			fmt.Println("AN ERROR OCCURRE WHILE READING INPUT")
			comm = -1
		}

		switch comm {
		case 0 | -1:
		case 1:
			var name string
			fmt.Print("Insert the node name: ")
			fmt.Scanf("%s", &name)
			g.AddNode(name)
		case 2:
			var a, b string
			var weight int
			fmt.Print("Insert the 2 node names separated by space: ")
			fmt.Scanf("%s %s", &a, &b)
			fmt.Print("Insert the weight: ")
			fmt.Scanf("%d", &weight)
			g.AddLine(a, b, weight)
		case 3:
			var a, b string
			fmt.Print("Insert the start and dest names separated by space: ")
			fmt.Scanf("%s %s", &a, &b)
			fmt.Printf("Distance: %d\n", g.BFS(a, b))
		case 4:
			var a, b string
			fmt.Print("Insert the start and dest names separated by space: ")
			fmt.Scanf("%s %s", &a, &b)
			fmt.Printf("Distance: %d\n", g.Dijkstra(a, b))
		}
	}
}