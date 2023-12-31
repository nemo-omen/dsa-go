package main

import (
	"dsa-go/list"
	"fmt"
)

func main() {
	list := list.New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)
	list.PushBack(7)
	list.PushBack(8)
	list.PushBack(9)
	list.PushBack(10)

	list.InsertBefore(101, 0)
	list.InsertBefore(66, 2)
	list.InsertAfter(96, 6)
	fmt.Println(list.ToSlice())
	list.RemoveByValue(101)
	list.RemoveByValue(66)
	list.RemoveByValue(96)
	fmt.Println(list.ToSlice())
	list.InsertAfter(999, list.Size()-1)
	fmt.Println(list.ToSlice())
	ninesNode := list.FindByValue(999)
	fmt.Println("ninesNode: ", ninesNode)
	list.Remove(ninesNode)
	fmt.Println(list.ToSlice())
}
