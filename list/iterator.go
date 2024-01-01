package list

type ListIterator[T comparable] struct {
	index int
	list  *List[T]
}

func (it *ListIterator[T]) HasNext() bool {
	if it.index < it.list.size {
		return true
	}
	return false
}

func (it *ListIterator[T]) Next() *Node[T] {
	if it.HasNext() {
		n := it.list.At(it.index)
		it.index += 1
		return n
	}
	return nil
}
