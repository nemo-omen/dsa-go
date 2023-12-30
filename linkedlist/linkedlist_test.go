package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	t.Run("Empty list should have nil head", func(t *testing.T) {
		l := New[int]()
		actual := l.Head
		assertNilPointer(t, actual)
	})

	t.Run("Empty list should have nil tail", func(t *testing.T) {
		l := New[int]()
		actual := l.Tail
		assertNilPointer(t, actual)
	})
}

func TestFront(t *testing.T) {
	t.Run("Front() should return zero value on new List", func(t *testing.T) {
		l := New[int]()
		actual := l.Front()
		expected := *new(int)
		assertEquals(t, expected, actual)
	})

	t.Run("Front() should return ListNode.Data on List with elements", func(t *testing.T) {
		l := New[int]()
		l.PushBack(1)
		l.PushBack(2)
		actual := l.Front()
		expected := 1
		assertEquals(t, expected, actual)
	})
}

func TestBack(t *testing.T) {
	t.Run("Back() should return T zero value on new List", func(t *testing.T) {
		l := New[int]()
		actual := l.Back()
		expected := *new(int)
		assertEquals(t, expected, actual)
	})

	t.Run("Back() should return ListNode.Data on List with elements", func(t *testing.T) {
		l := New[int]()
		l.PushBack(1)
		l.PushBack(2)
		actual := l.Back()
		expected := 2
		assertEquals(t, expected, actual)
	})
}

func TestAt(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	findTests := []struct {
		in       string
		list     *LinkedList[int]
		index    int
		expected any
	}{
		{"Test indexed access at 0", l, 0, 1},
		{"Test indexed access at 1", l, 1, 2},
		{"Test indexed access at 2", l, 2, 3},
		{"Test indexed access at 3", l, 3, 4},
		{"Test indexed access at 4", l, 4, 5},
		{"Test indexed access at 5", l, 5, fmt.Errorf("No such index %q", 5)},
	}

	for _, tt := range findTests {
		t.Run(tt.in, func(t *testing.T) {
			expected := tt.expected
			data, error := tt.list.At(tt.index)
			actual := data

			if error != nil {
				assertDeepEquals(t, expected, error)
			} else {
				assertEquals(t, expected, actual)
			}
		})
	}
}

func TestFind(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	atTests := []struct {
		in       string
		list     *LinkedList[int]
		value    int
		expected int
	}{
		{"Test find value 1", l, 1, 0},
		{"Test find value 2", l, 2, 1},
		{"Test find value 3", l, 3, 2},
		{"Test find value 4", l, 4, 3},
		{"Test find value 5", l, 5, 4},
		{"Test find value 5", l, 6, -1},
	}

	for _, tt := range atTests {
		t.Run(tt.in, func(t *testing.T) {
			expected := tt.expected
			actual := tt.list.Find(tt.value)

			assertEquals(t, expected, actual)
		})
	}
}

func TestPushFront(t *testing.T) {
	t.Run("PushFront() should create new ListNode and assign it to Head", func(t *testing.T) {
		l := New[int]()
		n := l.PushFront(1)
		expected := n
		actual := l.Head

		assertEquals(t, expected, actual)
	})

	t.Run("PushFront() on new List should assign new ListNode to Tail", func(t *testing.T) {
		l := New[int]()
		n := l.PushFront(1)
		expected := n
		actual := l.Tail
		assertEquals(t, expected, actual)
	})

	t.Run("PushFront() should increase the list's Size", func(t *testing.T) {
		l := New[int]()
		zero := l.Size
		l.PushFront(1)
		one := l.Size

		if !(one > zero) {
			t.Errorf("expected %d to be greater than %d", one, zero)
		}
	})
}

func TestPushBack(t *testing.T) {
	t.Run("PushBack() should create new ListNode and assign it to Tail", func(t *testing.T) {
		l := New[int]()
		n := l.PushBack(1)
		expected := n
		actual := l.Tail

		assertEquals(t, expected, actual)
	})

	t.Run("PushBack() on new List should assign new ListNode to Head", func(t *testing.T) {
		l := New[int]()
		n := l.PushBack(1)
		expected := n
		actual := l.Head
		assertEquals(t, expected, actual)
	})

	t.Run("PushBack() should increase the list's Size", func(t *testing.T) {
		l := New[int]()
		zero := l.Size
		l.PushBack(1)
		one := l.Size

		if !(one > zero) {
			t.Errorf("expected %d to be greater than %d", one, zero)
		}
	})
}

func TestPopFront(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	pbTests := []struct {
		in       string
		list     *LinkedList[int]
		expected any
	}{
		{"Test PopFront on 1st element", l, 1},
		{"Test PopFront on 2nd element", l, 2},
		{"Test PopFront on 3rd element", l, 3},
		{"Test PopFront on 4th element", l, 4},
		{"Test PopFront on 5th element", l, 5},
		{"Test PopFront on empty list", l, fmt.Errorf("Empty list cannot be popped")},
	}

	for _, tt := range pbTests {
		t.Run(tt.in, func(t *testing.T) {
			expected := tt.expected
			data, error := l.PopFront()
			actual := data

			if error != nil {
				assertDeepEquals(t, expected, error)
			} else {
				assertEquals(t, expected, actual)
			}
		})
	}
}

func TestPopBack(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	pbTests := []struct {
		in       string
		list     *LinkedList[int]
		expected any
	}{
		{"Test PopBack on 5th element", l, 5},
		{"Test PopBack on 4th element", l, 4},
		{"Test PopBack on 3rd element", l, 3},
		{"Test PopBack on 2nd element", l, 2},
		{"Test PopBack on 1st element", l, 1},
		{"Test PopBack on empty list", l, fmt.Errorf("Empty list cannot be popped")},
	}

	for _, tt := range pbTests {
		t.Run(tt.in, func(t *testing.T) {
			expected := tt.expected
			data, error := l.PopBack()
			actual := data

			if error != nil {
				assertDeepEquals(t, expected, error)
			} else {
				assertEquals(t, expected, actual)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	l := New[string]()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")
	l.PushBack("e")

	removeTests := []struct {
		in       string
		list     *LinkedList[string]
		value    string
		expected bool
	}{
		{"Test removal of value a", l, "a", true},
		{"Test removal of value b", l, "b", true},
		{"Test removal of value c", l, "c", true},
		{"Test removal of value e", l, "e", true},
		{"Test removal of value d", l, "d", true},
		{"Test removal of nonexistent value", l, "z", false},
	}

	for _, tt := range removeTests {
		t.Run(tt.in, func(t *testing.T) {
			expected := tt.expected
			actual := tt.list.Remove(tt.value)

			assertEquals(t, expected, actual)
		})
	}
}

func TestRemoveAt(t *testing.T) {
	removeTests := []struct {
		in       string
		index    int
		expected bool
	}{
		{"Test removal at index 0", 0, true},
		{"Test removal at index 1", 1, true},
		{"Test removal at index 2", 2, true},
		{"Test removal at index 3", 3, true},
		{"Test removal at index 4", 4, true},
		{"Test removal at index 5", 5, false},
	}

	for _, tt := range removeTests {
		l := New[string]()
		l.PushBack("a")
		l.PushBack("b")
		l.PushBack("c")
		l.PushBack("d")
		l.PushBack("e")

		t.Run(tt.in, func(t *testing.T) {
			expected := tt.expected
			actual := l.RemoveAt(tt.index)

			assertEquals(t, expected, actual)
		})
	}
}

func assertEquals(t testing.TB, expected, actual any) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func assertDeepEquals(t testing.TB, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func assertNilPointer(t testing.TB, actual any) {
	if !reflect.ValueOf(actual).IsNil() {
		t.Errorf("expected %v, actual %v", nil, actual)
	}
}

func assertNotNilPointer(t testing.TB, actual any) {
	if reflect.ValueOf(actual).IsNil() {
		t.Errorf("expected non-nil value, actual %v", actual)
	}
}
