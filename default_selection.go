package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond) // 100
	boom := time.After(5000 * time.Millisecond) // 500
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("      .")
			time.Sleep(1000 * time.Millisecond) // 50
		}
	}
}