package linkedlist

import (
	"fmt"
)

// LinkedList is a doubly-linked list
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

// Returns the value of the first ListNode in the list
func (l *LinkedList[T]) Front() T {
	if l.Head == nil {
		return *new(T)
	}
	return l.Head.Data
}

// Returns the value of the last ListNode in the list
func (l *LinkedList[T]) Back() T {
	if l.Tail == nil {
		return *new(T)
	}
	return l.Tail.Data
}

// Returns the value of the ListNode at the given index
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

// Returns index of first occurrence of a value in the list/
// Returns -1 if value not found. (Yep, just like C++ 😱)
func (l *LinkedList[T]) Find(value T) int {
	// TODO: Maybe you should create an iterator for this
	if l.Head == nil {
		return -1
	}

	current := l.Head
	index := 0

	for current != nil {
		if current.Data == value {
			return index
		}
		current = current.next
		index += 1
	}

	return -1
}

// Adds the given value to the beginning of the list
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

// Adds the given value to the end of the list
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

// Removes the first ListNode from the list and returns its value
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

// Removes the last ListNode from the list and returns its value
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

func (l *LinkedList[T]) RemoveAt(index int) bool {
	if index >= l.Size {
		return false
	}

	current := l.Head
	currentIndex := 0

	for currentIndex < index {
		current = current.next
		currentIndex += 1
	}

	if current == l.Head {
		if current.next != nil {
			current.next.previous = nil
			l.Head = current.next
		}
		current = &ListNode[T]{}
		l.Size -= 1
		return true
	} else if current == l.Tail {
		if current.previous != nil {
			current.previous.next = nil
			l.Tail = current.previous
		}
		current = &ListNode[T]{}
		l.Size -= 1
		return true
	} else if current != l.Head && current != l.Tail {
		if current.previous != nil && current.next != nil {
			current.previous.next = current.next
			current.next.previous = current.previous
		}
		current = &ListNode[T]{}
		l.Size -= 1
		return true
	} else {
		return false
	}
}

func (l *LinkedList[T]) Remove(value T) bool {
	index := l.Find(value)

	if index == -1 {
		return false
	}

	result := l.RemoveAt(index)
	return result
}
