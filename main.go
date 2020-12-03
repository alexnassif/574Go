package main

import (
	"./tree"
	"fmt"
	"sync"
	"time"
)

func Inorder(t *tree.Tree, wg *sync.WaitGroup) {
	defer wg.Done()
	var stack []*tree.Tree

	current := t

	for current != nil || len(stack) > 0 {

		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		n := len(stack) - 1
		current = stack[n]
		stack = stack[:n]
		current = current.Right
	}

}

func main() {
	var wg sync.WaitGroup
	// Code to measure
	t := tree.New()
	t1 := tree.New()

	start := time.Now()
	go Inorder(t, &wg)
	wg.Add(1)
	go Inorder(t1, &wg)
	wg.Add(1)

	wg.Wait()
	duration := time.Now()

	fmt.Println(duration)

	// Get duration.
	d := duration.Sub(start)
	fmt.Println("Duration", d)
	// Get seconds from duration.
	s := d.Seconds()
	fmt.Println("Seconds", s)

}
