package main

import (
	"basic"
	"fmt"
)

func main() {
	in := []int{1, 2, 3, 0, 0, 4, 0, 0, 5, 6, 0, 0, 7, 8, 8, 10, 0, 0}
	s := basic.NewArrayResizingQueue()
	for _, n := range in {
		if n != 0 {
			s.Enqueue(n)
		} else {
			i := s.Dequeue()
			fmt.Println(i)
		}
	}
}
