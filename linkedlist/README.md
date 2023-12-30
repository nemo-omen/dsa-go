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

## Notes

Counter to many educational examples (and `List` in Go's standard library), outside access to list elements will be done through the value stored within a `ListNode`'s `Data` property. Why? Because that's usually how one expects to interact with collections. For example, when a user is interacting with an array or slice, they don't expect to receive a struct like `ArrayElement` when they type `foo := myArray[5]`. They expect to receive the value contained at that index. I'll use `Head`, `Tail`, and each node's `Previous` and `Next` properties to access references to `ListNode`s internally, and other methods like `Front()`, `Back()`, and `At()` to access an element's `Data` for external use.

## Methods

- [x] `Front() T`
  - Returns the value held in the list's first element
- [x] `Back() T`
  - Returns the value held in the list's last element
- [x] `At(index int) (T, error)`
  - Returns the element at a given index
- [x] `Find(value T) (T, error)`
- [x] `PushFront(value T)`
  - Adds a value to the front of the list
- [x] `PushBack(value T)`
  - Adds a value to the back of the list
- [x] `InsertBefore(value T, index?) bool`
  - Inserts a value before a given index
- [x] `InsertAfter(value T, index?) bool`
  - Inserts a value after a given index
- [x] `PopBack() (T, error)`
  - Removes the last element in the list and returns the value of that element
- [x] `PopFront() (T, error)`
  - Removes the first element of the list and returns the value of that element
- [x] `Remove(value T) (T, error)`
- [x] `RemoveAt(index int) (T, error)`
- [x] `ToSlice() []T`

## Potential Changes

### 1. Use a sentinel root node.
As described in [Introduction to Algorithms](https://www.amazon.com/Introduction-Algorithms-fourth-Thomas-Cormen/dp/026204630X?keywords=introduction+to+algorithms&qid=1660584469&s=books&sprefix=introduction,stripbooks,93&sr=1-1&linkCode=sl1&tag=theprimeagen-20&linkId=83c55b2c018fbe94a1de79ebd447b66a&language=en_US&ref_=as_li_ss_tl), and as seen in the [std library implementation](https://cs.opensource.google/go/go/+/refs/tags/go1.21.5:src/container/list/list.go) of `List`, the list can be made circular with the use of a `nil` root Node. This removes the need for the list to hold a reference to `Head` and `Tail`.

### 2. Rethink insertion and removal code
The code works, but it could probably be a little more concise. Maybe do some more refactoring to clean it up.

### 3. Initialize with Slice/Array
You should be able to initialize a `LinkedList` from an array or slice.