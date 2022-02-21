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
		fmt.Println("4 - DFS")
		fmt.Println("5 - Dijkstra")
		fmt.Println()

		var com string
		var err error

		fmt.Print("> ")
		fmt.Scanf("%s", &com)

		comm, err = strconv.Atoi(com)
		if err != nil {
			fmt.Println("AN ERROR OCCURRED WHILE READING INPUT")
			comm = -1
		}

		switch comm {
		case 0 | -1:
		case 1:
			AddNode(&g)
		case 2:
			AddLine(&g)
		case 3:
			BFS(g)
		case 4:
			DFS(g)
		case 5:
			Dijkstra(g)
		}
	}
}

func AddNode(g *graphs.Graph) {
	var name string
	fmt.Print("Insert the node name: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Printf("an error occured while reading input")
		return
	}
	g.AddNode(name)
}

func AddLine(g *graphs.Graph) {
	var a, b string
	var weight int

	fmt.Print("Insert the 2 node names separated by space: ")
	_, err := fmt.Scanf("%s %s", &a, &b)
	if err != nil {
		fmt.Printf("an error occured while reading input")
		return
	}

	fmt.Print("Insert the weight: ")
	_, err = fmt.Scanf("%d", &weight)
	if err != nil {
		fmt.Printf("an error occured while reading input")
		return
	}

	g.AddLine(a, b, weight)
}

func BFS(g graphs.Graph) {
	var a, b string

	fmt.Print("Insert the start and dest names separated by space: ")
	_, err := fmt.Scanf("%s %s", &a, &b)

	if err != nil {
		fmt.Printf("an error occured while reading input")
		return
	}

	fmt.Printf("Distance: %d\n", g.BFS(a, b))
}

func Dijkstra(g graphs.Graph) {
	var a, b string

	fmt.Print("Insert the start and dest names separated by space: ")
	_, err := fmt.Scanf("%s %s", &a, &b)

	if err != nil {
		fmt.Printf("an error occured while reading input")
		return
	}

	cost, path := g.Dijkstra(a, b)
	fmt.Printf("Path: %s\nCost: %d\n\n", path, cost)
}

func DFS(g graphs.Graph) {
	var a string

	fmt.Print("Insert the start node name: ")
	_, err := fmt.Scanf("%s", &a)

	if err != nil {
		fmt.Printf("an error occured while reading input")
		return
	}

	fmt.Printf("Not connected nodes: %s\n", g.DFS(a))
}
