package linkedlist

type Iterator[T comparable] interface {
	hasNext() bool
	next() *ListNode[T]
}

type LinkedListIterator[T comparable] struct {
	index int
	list  *LinkedList[T]
}

func (li *LinkedListIterator[T]) hasNext() bool {
	return false
}

func (li *LinkedListIterator[T]) next() *ListNode[T] {
	return &ListNode[T]{}
}
