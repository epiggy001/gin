// Copyright 2015 Clustertech Limited. All rights reserved.
//
// Author: jackeychen (jackeychen@clustertech.com)
package main

import (
	"fmt"
	"math/rand"
	"sorter"
)

const kTestTime = 100

func isSorted(a []int) bool {
	l := len(a)
	for i := 0; i < l-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func initArray(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}
	return a
}

func testSort(name string, f func([]int)) {
	for i := 0; i < kTestTime; i++ {
		a := initArray(20)
		f(a)
		if !isSorted(a) {
			fmt.Println(name, ": Fail")
			return
		}
	}
	fmt.Println(name, ": Success")

}

func main() {
	testSort("Insertion", sorter.Insertion)
	testSort("Selection", sorter.Selection)
	testSort("Shell", sorter.Shell)
	testSort("Merge", sorter.Merge)
}
