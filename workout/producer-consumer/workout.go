package main

import (
	"fmt"
	"time"
)

type store struct {
	q chan int
}

func main() {
	s := store{q: make(chan int)}

	go func() {
		counter := 0
		for {
			fmt.Printf("producing: %d\n", counter)
			s.q <- counter
			counter += 1
		}
	}()

	for counter := range s.q {
		fmt.Printf("consuming: %d\n", counter)
		time.Sleep(1 * time.Second)
	}
}
