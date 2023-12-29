package main

import (
	"dsa-go/linkedlist"
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	jeff := Person{
		name: "Jeff",
		age:  46,
	}

	list := linkedlist.New[Person]()
	list.PushBack(jeff)
	fmt.Println("List head: ", list.Head)
}
