package main

import (
	"./tree"
	"fmt"
	"time"
)


func Inorder(t *tree.Tree, ch chan int) {
	var stack []*tree.Tree

	current := t

	for current != nil || len(stack) > 0{

		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		n := len(stack) - 1
		current = stack[n]
		stack = stack[:n]
		ch <- current.Value
		current = current.Right
	}

	close(ch)
}



func main(){
	times := make([] float64, 10)
	for i := 0; i < 10; i++ {
		// Code to measure
		t := tree.New()
		t1 := tree.New()
		ch := make(chan int)
		ch1 := make(chan int)

		start := time.Now()
		go Inorder(t, ch)
		go Inorder(t1, ch1)

		for i := 0; i < tree.SIZE; i++ {
			fmt.Print("ch ")
			fmt.Println(<-ch)
		}
		for i := 0; i < tree.SIZE; i++ {
			fmt.Print("ch1 ")
			fmt.Println(<-ch1)
		}
		duration := time.Now()

		fmt.Println(duration)

		// Get duration.
		d := duration.Sub(start)
		fmt.Println("Duration", d)
		// Get seconds from duration.
		s := d.Seconds()
		fmt.Println("Seconds", s)
		times = append(times, s)

	}
	sumOfTimes := 0.0
	for _, num := range times {
		sumOfTimes += num
	}

	fmt.Println(sumOfTimes/10)


}