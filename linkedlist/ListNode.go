package linkedlist

import "reflect"

type ListNode[T comparable] struct {
	previous, next *ListNode[T]
	Data           T
}

func (n *ListNode[T]) NewNode(value T) ListNode[T] {
	return ListNode[T]{Data: value, previous: nil, next: nil}
}

// Returns the given node's Next node or nil
func (n *ListNode[T]) Next() *ListNode[T] {
	if !reflect.DeepEqual(n.next, nil) {
		return n.next
	}
	return nil
}

// Returns the given node's previous node or nil
func (n *ListNode[T]) Previous() *ListNode[T] {
	if !reflect.DeepEqual(n.previous, nil) {
		return n.previous
	}
	return nil
}
