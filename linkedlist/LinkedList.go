package linkedlist

import (
	"fmt"
)

type LinkedList[T comparable] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
	Size int
}

// Initializes a new LinkedList with nil Head,
// nil Tail, and 0 Size
func (l *LinkedList[T]) Init() *LinkedList[T] {
	l.Head = nil
	l.Tail = nil
	l.Size = 0
	return l
}

// Returns an initialized LinkedList
func New[T comparable]() *LinkedList[T] {
	return new(LinkedList[T]).Init()
}

func (l *LinkedList[T]) Front() T {
	if l.Head == nil {
		return *new(T)
	}
	return l.Head.Data
}

func (l *LinkedList[T]) Back() T {
	if l.Tail == nil {
		return *new(T)
	}
	return l.Tail.Data
}

func (l *LinkedList[T]) At(index int) (T, error) {
	if index == 0 {
		return l.Head.Data, nil
	}

	if index == l.Size-1 {
		return l.Tail.Data, nil
	}

	if index >= l.Size {
		return *new(T), fmt.Errorf("No such index %q", index)
	}

	count := 0
	current := l.Head

	for count < index {
		current = current.next
		count += 1
	}
	return current.Data, nil
}

func (l *LinkedList[T]) PushBack(value T) *ListNode[T] {
	n := ListNode[T]{nil, nil, value}

	if l.Tail == nil {
		l.Head = &n
		l.Tail = &n
	} else {
		n.previous = l.Tail
		l.Tail.next = &n
		l.Tail = &n
	}
	l.Size += 1
	return l.Tail
}

func (l *LinkedList[T]) PushFront(value T) *ListNode[T] {
	n := ListNode[T]{nil, nil, value}

	if l.Head == nil {
		l.Head = &n
		l.Tail = &n
	} else {
		n.next = l.Head
		l.Head.previous = &n
		l.Head = &n
	}
	l.Size += 1
	return l.Head
}

func (l *LinkedList[T]) PopBack() (T, error) {
	if l.Tail == nil {
		return *new(T), fmt.Errorf("Empty list cannot be popped")
	}

	n := l.Tail
	data := n.Data

	if l.Tail == l.Head {
		l.Head = nil
		l.Tail = nil
	} else {
		n.previous.next = nil
		l.Tail = n.previous
	}

	n = &ListNode[T]{}
	l.Size -= 1
	return data, nil
}

func (l *LinkedList[T]) PopFront() (T, error) {
	if l.Head == nil {
		return *new(T), fmt.Errorf("Empty list cannot be popped")
	}

	n := l.Head
	data := n.Data

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		n.next.previous = nil
		l.Head = n.next
	}

	n = &ListNode[T]{}
	l.Size -= 1
	return data, nil
}
