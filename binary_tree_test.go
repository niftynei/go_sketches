package main

import (
	"golang.org/x/tour/tree"
	"sort"
	"testing"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	trees := []int{}
	for i := 0; i < 10; i++ {
		val := <-ch
		trees = append(trees, val)
	}

	sort.Ints(trees)
	for i := 0; i < 10; i++ {
		if i+1 != trees[i] {
			t.Errorf("Invalid match: expected %v != actual %v", i+1, trees[i])
		}
	}
}

func TestSameWorks(t *testing.T) {
	if !Same(tree.New(1), tree.New(1)) {
		t.Errorf("Trees are not the same, but should be")
	}
}

func TestSameNoWork(t *testing.T) {
	if Same(tree.New(2), tree.New(1)) {
		t.Errorf("Trees are the same, but should not be")
	}
}
