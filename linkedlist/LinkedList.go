package linkedlist

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

func (l *LinkedList[T]) Front() *ListNode[T] {
	return l.Head
}

func (l *LinkedList[T]) Back() *ListNode[T] {
	return l.Tail
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
