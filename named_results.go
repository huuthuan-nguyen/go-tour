package main

import "fmt"

func split(sum int) (x, y int) {
	z := 9
	x = sum * 4 / z
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}