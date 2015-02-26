// Copyright 2015 Clustertech Limited. All rights reserved.
//
// Author: jackeychen (jackeychen@clustertech.com)
package main

import (
	"basic"
	"fmt"
)

func main() {
	in := []int{1, 2, 3, 0, 0, 4, 0, 0, 5, 6, 0, 0, 7, 8, 8, 10, 0, 0}
	s := basic.NewArrayResizingStack()
	for _, n := range in {
		if n != 0 {
			s.Push(n)
		} else {
			i, err := s.Pop()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(i)
		}
	}
}
