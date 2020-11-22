package main

import (
	"./tree"
	"fmt"
	"time"
)

func Inorder(t *tree.Tree, ch chan int) {
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
		ch <- current.Value
		current = current.Right
	}

	close(ch)
}

func main() {
	messages := make(chan string)
	go world(messages)
	go hello(messages)
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(<- messages)
}
func hello(messages chan string){
	hello := "Hello"
	messages <- hello
}
func world(messages chan string)  {
	world := <-messages
	world += " World"
	messages <- world
}
