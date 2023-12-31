package list

type Node[T comparable] struct {
	Data     T
	Previous *Node[T]
	Next     *Node[T]
}
