package list

import "reflect"

type List[T comparable] struct {
	root *Node[T]
	size int
}

// / Initializes a list with a sentinel node
// / This is technically a circular, doubly-linked
// / list. The root node contains no Data and serves
// / to point to the first and last elements.
func (l *List[T]) Init() *List[T] {
	l.root = &Node[T]{}
	l.root.Next = l.root
	l.root.Previous = l.root
	l.size = 0
	return l
}

// / Initialize and return an initialized List
func New[T comparable]() *List[T] {
	return new(List[T]).Init()
}

// / Create a list from a slice
func NewFromSlice[T comparable](slice []T) *List[T] {
	l := new(List[T]).Init()

	for _, el := range slice {
		l.PushBack(el)
	}
	return l
}

/// Create an iterator for the List

func (l *List[T]) CreateIterator() *ListIterator[T] {
	return &ListIterator[T]{0, l}
}

// / Returns the number of elements in the list
func (l *List[T]) Size() int {
	return l.size
}

// ELEMENT ACCESS //

// / Returns the first Node in the list
func (l *List[T]) Front() *Node[T] {
	if l.size == 0 {
		return nil
	}
	return l.root.Next
}

// / Returns the last Node in the list
func (l *List[T]) Back() *Node[T] {
	if l.size == 0 {
		return nil
	}
	return l.root.Previous
}

// / Move "cursor" to an element at the given index.
// / Iterates forward from beginning of list.
func (l *List[T]) to(index int) *Node[T] {
	currentIndex := 0
	currentNode := l.root.Next
	for currentIndex < index {
		currentNode = currentNode.Next
		currentIndex += 1
	}
	if currentNode == l.root {
		return nil
	}
	return currentNode
}

// / Move "cursor" to an element at a given index/
// / Iterates backward from end of list.
func (l *List[T]) reverseTo(index int) *Node[T] {
	currentIndex := l.size - 1
	currentNode := l.root.Previous
	for currentIndex > index {
		currentNode = currentNode.Previous
		currentIndex -= 1
	}
	if currentNode == l.root {
		return nil
	}
	return currentNode
}

// / Returns the element at the given index
// / Can handle negative indices
func (l *List[T]) At(index int) *Node[T] {
	if index >= l.size {
		return nil
	}

	if index < 0 {
		// Handle negative indices by adding them
		// to size - 1
		rindex := (l.size - 1) + index
		if rindex < 0 {
			// if index is still negative
			// don't loop (don't want to lap)
			return nil
		}
		return l.At(rindex)
	}

	var node *Node[T]

	if l.size-index < index {
		// If our index is closer to the back than
		// the front, traverse backward
		node = l.reverseTo(index)
	} else {
		// Otherwise, traverse forward
		node = l.to(index)
	}
	return node
}

// / Finds a Node according to its value.
// / Resuturns first occurrence of given value.
func (l *List[T]) FindByValue(value T) *Node[T] {
	currentNode := l.root.Next
	for currentNode != l.root {
		if currentNode.Data == value {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}

// / Returns the index of the given Node or -1 if not found.
func (l *List[T]) FindIndex(node *Node[T]) int {
	currentNode := l.root.Next
	currentIndex := 0

	for currentNode != l.root {
		if currentNode == node {
			return currentIndex
		}
		currentNode = currentNode.Next
		currentIndex += 1
	}
	return -1
}

// ELEMENT INSERTION //

// / Adds value to back of list.
func (l *List[T]) PushBack(v T) *Node[T] {
	node := Node[T]{v, nil, nil}
	if !reflect.DeepEqual(l.root, l.root.Previous) {
		l.root.Previous.Next = &node
	} else {
		l.root.Next = &node
	}
	node.Previous = l.root.Previous
	l.root.Previous = &node
	node.Next = l.root
	l.size += 1
	return &node
}

// / Adds value to front of list.
func (l *List[T]) PushFront(v T) *Node[T] {
	node := Node[T]{v, nil, nil}
	if !reflect.DeepEqual(l.root, l.root.Next) {
		l.root.Next.Previous = &node
	} else {
		l.root.Previous = &node
	}
	node.Next = l.root.Next
	l.root.Next = &node
	node.Previous = l.root
	l.size += 1
	return &node
}

// / Inserts value before a given index
// / Returns a Node with the given value as Node.Data
func (l *List[T]) InsertBefore(value T, index int) *Node[T] {
	if index == 0 {
		return l.PushFront(value)
	}

	node := Node[T]{value, nil, nil}
	target := l.At(index)
	if target != nil {
		target.Previous.Next = &node
		node.Previous = target.Previous
		target.Previous = &node
		node.Next = target
		l.size += 1
		return &node
	}
	return nil
}

// / Inserts a value after the given index
// / Returns a Node with the value as Node.Data
func (l *List[T]) InsertAfter(value T, index int) *Node[T] {
	if index == l.size-1 {
		return l.PushBack(value)
	}

	node := Node[T]{value, nil, nil}
	target := l.At(index)
	if target != nil {
		target.Next.Previous = &node
		node.Next = target.Next
		target.Next = &node
		node.Previous = target
		l.size += 1
		return &node
	}
	return nil
}

// ELEMENT REMOVAL //

// / Removes a Node from the list
// / returns true if successful, false if failed
func (l *List[T]) Remove(node *Node[T]) bool {
	if node == l.root {
		return false
	}
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
	return true
}

// / Removes the node with the given value
// / Returns true if successful, false if failed
func (l *List[T]) RemoveByValue(value T) bool {
	node := l.FindByValue(value)
	return l.Remove(node)
}

// / Removes a Node at the given index
// / Returns true if successful, false if failed
func (l *List[T]) RemoveAt(index int) bool {
	node := l.At(index)
	return l.Remove(node)
}

// / Removes Node at end of list
// / returns that Node if successful
// / or nil if unsuccessful
func (l *List[T]) PopBack() *Node[T] {
	node := l.Back()
	result := l.Remove(node)
	if result == true {
		return node
	}
	return nil
}

// / Removes Node at beginning of list
// / returns that Node if successful
// / or nil if unsuccessful
func (l *List[T]) PopFront() *Node[T] {
	node := l.Front()
	result := l.Remove(node)
	if result == true {
		return node
	}
	return nil
}

// UTIL //

// / Returns a slice representation of the list
func (l *List[T]) ToSlice() *[]T {
	slice := []T{}
	current := l.root.Next
	for current != l.root {
		slice = append(slice, current.Data)
		current = current.Next
	}
	return &slice
}

// / Runs a given function over all elements of the list
// / Function must return the same type as the list
// / Returns a new list.
func (l *List[T]) Map(f func(T) T) *List[T] {
	newList := New[T]()
	currentNode := l.Front()

	for currentNode != l.root {
		data := currentNode.Data
		transformed := f(data)
		newList.PushBack(transformed)
		currentNode = currentNode.Next
	}
	return newList
}
