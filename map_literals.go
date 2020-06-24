package main

import "fmt"

type Vertex struct {
	Lat, Lng float64
}

var m = map[string]Vertex{
	"Home": Vertex{
		40.68433, -74.39967,
	},
	"Office": Vertex{
		41.99121, -80.34211,
	},
}

func main() {
	fmt.Println(m)
}