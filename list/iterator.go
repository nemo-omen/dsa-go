package list

type ListIterator[T comparable] struct {
	index int
	list  *List[T]
}

func (it *ListIterator[T]) hasNext() bool {
	if it.index < it.list.size {
		return true
	}
	return false
}

func (it *ListIterator[T]) next() *Node[T] {
	if it.hasNext() {
		n := it.list.At(it.index)
		it.index += 1
		return n
	}
	return nil
}
