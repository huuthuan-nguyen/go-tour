package main

import "fmt"

// interface with M method.
type I interface {
	M()
}

// T struct with S string.
type T struct {
	S string
}

// implement the M method for T type.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var t T = T{"Hello"}
	t.M()
}