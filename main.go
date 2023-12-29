package main

import (
	"fmt"

	"dsa-go/linkedlist"
)

func main() {
	list := linkedlist.New[int]()

	n := list.PushBack(1)
	fmt.Println("list: ", list)
	fmt.Println("n before: ", n)
	n2 := list.PushBack(2)
	fmt.Println("n after: ", n)
	fmt.Println("n2: ", n2)
	fmt.Println("list: ", list)
}
