package list

type Mapper[T comparable, M comparable] interface {
	Map()
}

func Map[T comparable, M comparable](l *List[T], f func(T) M) *List[M] {
	newlist := New[M]()
	iterator := l.CreateIterator()

	for iterator.HasNext() {
		val := iterator.Next().Data
		transformed := f(val)
		newlist.PushBack(transformed)
	}
	return newlist
}
