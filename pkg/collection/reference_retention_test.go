package collection

import "testing"

func TestQueueRemoveZeroesReleasedSlot(t *testing.T) {
	t.Parallel()

	queue := NewQueueWithCapacity[*int](4)
	first := intPtr(1)
	second := intPtr(2)
	queue.Add(first)
	queue.Add(second)

	before := queue.elements

	removed, ok := queue.Remove()
	if !ok {
		t.Fatal("Remove() returned ok=false")
	}
	if removed != first {
		t.Fatalf("Remove() removed %p, want %p", removed, first)
	}
	if before[0] != nil {
		t.Fatalf("released queue slot retained %p, want nil", before[0])
	}
}

func TestQueueClearZeroesReleasedSlots(t *testing.T) {
	t.Parallel()

	queue := NewQueueWithCapacity[*int](4)
	queue.Add(intPtr(1))
	queue.Add(intPtr(2))

	before := queue.elements

	queue.Clear()

	if queue.Len() != 0 {
		t.Fatalf("Len() after Clear() = %d, want 0", queue.Len())
	}
	for index := range before {
		if before[index] != nil {
			t.Fatalf("queue slot %d retained %p, want nil", index, before[index])
		}
	}
}

func TestStackPopZeroesReleasedSlot(t *testing.T) {
	t.Parallel()

	stack := NewStackWithCapacity[*int](4)
	first := intPtr(1)
	second := intPtr(2)
	stack.Push(first)
	stack.Push(second)

	before := stack.elements

	popped, ok := stack.Pop()
	if !ok {
		t.Fatal("Pop() returned ok=false")
	}
	if popped != second {
		t.Fatalf("Pop() returned %p, want %p", popped, second)
	}
	if before[len(before)-1] != nil {
		t.Fatalf("released stack slot retained %p, want nil", before[len(before)-1])
	}
}

func TestStackClearZeroesReleasedSlots(t *testing.T) {
	t.Parallel()

	stack := NewStackWithCapacity[*int](4)
	stack.Push(intPtr(1))
	stack.Push(intPtr(2))

	before := stack.elements

	stack.Clear()

	if stack.Len() != 0 {
		t.Fatalf("Len() after Clear() = %d, want 0", stack.Len())
	}
	for index := range before {
		if before[index] != nil {
			t.Fatalf("stack slot %d retained %p, want nil", index, before[index])
		}
	}
}

func intPtr(value int) *int {
	return &value
}
