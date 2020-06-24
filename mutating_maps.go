package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Printf("The value's %d\n", m["Answer"])

	m["Answer"] = 48
	fmt.Printf("The value's %d\n", m["Answer"])

	delete(m, "Answer")
	fmt.Printf("The value's %d\n", m["Answer"])

	v, ok := m["Answer"]
	fmt.Printf("The value's %d Present? %t", v, ok)
}