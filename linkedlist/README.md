# Linked List

A linked list is similar to an array in that it can be thought of as a linear collection of elements. Unlike arrays in most languages (not JavaScript), linked lists are stored in heap memory.

In this implementation, the list is represented by a `LinkedList` type, which holds references to the first element, `Head`, the last element, `Tail`, and the number of elements in the container, `Size`.

The `ListNode` type contains whatever data is being stored within the node, `Data`, and references to the `previous` and `next` nodes.

```go
type ListNode[T] struct {
  Data     T
  previous *ListNode[T]
  next     *ListNode[T]
}

type LinkedList[T] {
  Head *ListNode[T]
  Tail *ListNode[T]
  Size int
}
```

## Implementation Notes

Counter to many educational examples (and `List` in Go's standard library), outside access to list elements will be done through the value stored within a `ListNode`'s `Data` property. Why? Because that's usually how one expects to interact with collections. For example, when a user is interacting with an array or slice, they don't expect to receive a struct like `ArrayElement` when they type `foo := myArray[5]`. They expect to receive the value contained at that index. I'll use `Head`, `Tail`, and each node's `Previous` and `Next` properties to access references to `ListNode`s internally, and other methods like `Front()`, `Back()`, and `At()` to access an element's `Data` for external use.

## Methods

- [x] `Front() T`
- [x] `Back() T`
- [x] `At(index int) (T, error)`
- [x] `PushFront(value T)`
- [x] `PushBack(value T)`
- [ ] `PopBack() (T, error)`
- [ ] `PopFront() (T, error)`
- [ ] `Remove(value T) (T, error)`
- [ ] `InsertBefore(value T, index?) (T, error)`
- [ ] `InsertAfter(value T, index?) (T, error)`
- [ ] `Has(value T) boolean`
- [ ] `Find(value T) (T, error)`