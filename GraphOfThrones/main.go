package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mygraph := IngestInput()
	if mygraph.Balance() {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not balanced")
	}
}

// IngestInput ingest input
func IngestInput() Graph {
	mygraph := Graph{
		nodes:  []*Node{},
		enemy:  map[Node][]*Node{},
		friend: map[Node][]*Node{},
	}
	reader := bufio.NewReader(os.Stdin)
	l, _ := reader.ReadString('\n')
	l = strings.Replace(l, "\n", "", 1)
	cl := strings.Split(l, " ")
	ltr, _ := strconv.Atoi(cl[1])
	for i := 0; i < ltr; i++ {
		l, _ := reader.ReadString('\n')
		l = strings.Replace(l, "\n", "", 1)
		friends := strings.Split(l, " ++ ")
		enemy := strings.Split(l, " -- ")
		if len(friends) == 2 {
			c1 := mygraph.NewNode(friends[0])
			c2 := mygraph.NewNode(friends[1])
			mygraph.AddEdge(c1, c2, true)
		} else {
			c1 := mygraph.NewNode(enemy[0])
			c2 := mygraph.NewNode(enemy[1])
			mygraph.AddEdge(c1, c2, false)
		}
	}
	return mygraph
}
