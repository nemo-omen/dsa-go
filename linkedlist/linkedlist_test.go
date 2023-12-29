package linkedlist

import (
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
}

func TestFront(t *testing.T) {
	t.Run("Front() should return nil on new List", func(t *testing.T) {
		l := New[int]()
		actual := l.Front()

		assertNilPointer(t, actual)
	})

	t.Run("Front() should return ListNode on List with elements", func(t *testing.T) {
		l := New[int]()
		n := l.PushBack(1)
		l.PushBack(2)
		actual := l.Front()
		assertEquals(t, n, actual)
		// assertNotNilPointer(t, actual)
	})
}

func TestBack(t *testing.T) {
	t.Run("Back() should return nil on new List", func(t *testing.T) {
		l := New[int]()
		actual := l.Back()
		assertNilPointer(t, actual)
	})

	t.Run("Back() should return ListNode on List with elements", func(t *testing.T) {
		l := New[int]()
		l.PushBack(1)
		n := l.PushBack(2)
		actual := l.Back()
		assertEquals(t, n, actual)
	})
}

func assertEquals(t testing.TB, expected, actual any) {
	t.Helper()
	if expected != actual {
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
