package main

import (
	"fmt"
	"math"
)

// define the interface.
type Abser interface {
	Abs() float64
}

type MyFloat float64

// MyFloat implement the Abser.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// *Vertex implement the Abser.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	var v *Vertex = &Vertex{3, 4}
	
	f := MyFloat(-math.Sqrt2)

	a = f // MyFloat implement Abser.
	// a = &v // *Vertex implement Abser.

	a = v // Vertex does NOT implement Abser, just *Vertex is implemented the Abser with Abs method.
	fmt.Println(a.Abs())
}