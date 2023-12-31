package collection

type Collection[T comparable] interface {
	createIterator() Iterator[T]
}
