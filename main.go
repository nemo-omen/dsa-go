package main

import (
	"dsa-go/list"
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	list1 := list.NewFromSlice[int](s)
	newList := list1.Map(func(n int) int {
		return n * 2
	})
	fmt.Println(newList.ToSlice())
	newList2 := mapList[int, string](newList, func(i int) string {
		return fmt.Sprintf("n%d", i)
	})
	fmt.Println(newList2.ToSlice())
}

func mapList[T comparable, M comparable](l *list.List[T], f func(T) M) *list.List[M] {
	newlist := list.New[M]()
	iterator := l.CreateIterator()

	for iterator.HasNext() {
		val := iterator.Next().Data
		transformed := f(val)
		newlist.PushBack(transformed)
	}
	return newlist
}
