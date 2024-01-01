package list

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {

	t.Run("Test new list root.Next points to root", func(t *testing.T) {
		l := New[int]()
		expected := l.root
		actual := l.root.Next

		assertDeepEquals(t, expected, actual)
	})

	t.Run("Test new list root.Previous points to root", func(t *testing.T) {
		l := New[int]()
		expected := l.root
		actual := l.root.Previous

		assertDeepEquals(t, expected, actual)
	})
}

func TestNewFromSlice(t *testing.T) {
	type TestStruct struct {
		whatever  string
		thatThing int
	}

	l1 := New[int]()
	l2 := New[string]()
	l3 := New[TestStruct]()

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	s1 := []int{1, 2, 3}

	l2.PushBack("a")
	l2.PushBack("b")
	l2.PushBack("c")
	s2 := []string{"a", "b", "c"}

	l3.PushBack(TestStruct{"no", 0})
	l3.PushBack(TestStruct{"yes", 1})
	l3.PushBack(TestStruct{"maybe", -1})
	s3 := []TestStruct{{"no", 0}, {"yes", 1}, {"maybe", -1}}

	t.Run("Test NewFromSlice on slice of ints", func(t *testing.T) {
		expected := l1
		actual := NewFromSlice[int](s1)
		assertDeepEquals(t, expected, actual)
	})

	t.Run("Test NewFromSlice on slice of strings", func(t *testing.T) {
		expected := l2
		actual := NewFromSlice[string](s2)
		assertDeepEquals(t, expected, actual)
	})

	t.Run("Test NewFromSlice on slice of structs", func(t *testing.T) {
		expected := l3
		actual := NewFromSlice[TestStruct](s3)
		assertDeepEquals(t, expected, actual)
	})
}

func TestSingleNode(t *testing.T) {
	t.Run("Test Node on single-node list points to l.root as previous", func(t *testing.T) {
		l := New[int]()
		l.PushBack(1)
		expected := l.root
		actual := l.Front().Previous
		assertDeepEquals(t, expected, actual)
	})

	t.Run("Test Node on single-node list points to l.root as next", func(t *testing.T) {
		l := New[int]()
		l.PushBack(1)
		expected := l.root
		actual := l.Front().Next
		assertDeepEquals(t, expected, actual)
	})

	t.Run("Test l.root.Next points to first node", func(t *testing.T) {
		l := New[int]()
		n := l.PushBack(1)
		expected := n
		actual := l.root.Next
		assertDeepEquals(t, expected, actual)
	})

	t.Run("Test l.root.Previous points to first node", func(t *testing.T) {
		l := New[int]()
		n := l.PushBack(1)
		expected := n
		actual := l.root.Previous
		assertDeepEquals(t, expected, actual)
	})
}

func TestFront(t *testing.T) {
	t.Run("Test Front() on new list returns nil", func(t *testing.T) {
		l := New[int]()
		actual := l.Front()
		assertNilPointer(t, actual)
	})

	t.Run("Test Front() returns Node", func(t *testing.T) {
		l := New[int]()
		l.PushBack(1)
		actual := l.Front()
		expected := &Node[int]{1, l.root, l.root}
		assertDeepEquals(t, expected, actual)
	})
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
