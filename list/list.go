package list

type List[T comparable] struct {
	root *Node[T]
	size int
}

func (l *List[T]) Init() *List[T] {
	l.root = &Node[T]{}
	l.root.Next = l.root
	l.root.Previous = l.root
	l.size = 0
	return l
}

func New[T comparable]() *List[T] {
	return new(List[T]).Init()
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Front() *Node[T] {
	if l.size == 0 {
		return nil
	}
	return l.root.Next
}

func (l *List[T]) Back() *Node[T] {
	if l.size == 0 {
		return nil
	}
	return l.root.Previous
}

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
			return nil
		}
		return l.At(rindex)
	}

	var node *Node[T]

	if l.size-index < index {
		// If our index is closer to the back than
		// the front, traverse backward
		// NOTE: I think this reduces time complexity
		// to O(log n) by reducing input size by half
		node = l.reverseTo(index)
	} else {
		// Otherwise, traverse forward
		node = l.to(index)
	}
	return node
}

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

func (l *List[T]) PushBack(v T) *Node[T] {
	node := Node[T]{v, nil, nil}
	if l.root.Previous != nil {
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

func (l *List[T]) PushFront(v T) *Node[T] {
	node := Node[T]{v, nil, nil}
	if l.root.Next != nil {
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

func (l *List[T]) Remove(node *Node[T]) bool {
	if node == nil {
		return false
	}
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
	return true
}

func (l *List[T]) RemoveByValue(value T) bool {
	node := l.FindByValue(value)
	return l.Remove(node)
}

func (l *List[T]) RemoveAt(index int) bool {
	node := l.At(index)
	return l.Remove(node)
}

func (l *List[T]) PopBack() *Node[T] {
	node := l.root.Previous
	l.root.Previous = node.Previous
	node.Previous.Next = l.root
	return node
}

func (l *List[T]) PopFront() *Node[T] {
	node := l.root.Next
	l.root.Next = node.Next
	node.Next.Previous = l.root
	return node
}

func (l *List[T]) ToSlice() *[]T {
	slice := []T{}
	current := l.root.Next
	for current != l.root {
		slice = append(slice, current.Data)
		current = current.Next
	}
	return &slice
}