package main

import (
	"golang.org/x/tour/tree"
	"sort"
)

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil && *t.Left != (tree.Tree{}) {
		Walk(t.Left, ch)
	}
	if t.Right != nil && *t.Right != (tree.Tree{}) {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	values1, values2 := []int{}, []int{}
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		val1 := <-ch1
		val2 := <-ch2
		values1 = append(values1, val1)
		values2 = append(values2, val2)
	}

	sort.Ints(values1)
	sort.Ints(values2)

	for i := range values1 {
		if values1[i] != values2[i] {
			return false
		}
	}
	return true
}
