package collection

// Queue is a generic non thread-safe unbounded FIFO queue implementation.
type Queue[T any] struct {
	elements []T
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// NewQueueWithCapacity creates a new queue with the specified initial capacity.
func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		elements: make([]T, 0, capacity),
	}
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.elements = q.elements[:0]
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return len(q.elements)
}

// Cap returns the capacity of the queue.
func (q *Queue[T]) Cap() int {
	return cap(q.elements)
}

// Add adds an element to the queue.
func (q *Queue[T]) Add(element T) {
	q.elements = append(q.elements, element)
}

// Peek returns the next element to be removed from the queue (without removing it)
// and a boolean indicating if the operation was successful.
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.elements) == 0 {
		var result T
		return result, false
	}

	return q.elements[0], true
}

// Remove removes and returns the next element of the queue and a boolean indicating
// if the operation was successful.
func (q *Queue[T]) Remove() (T, bool) {
	if len(q.elements) == 0 {
		var result T
		return result, false
	}

	result := q.elements[0]
	// TODO: Confirm whether this is necessary to prevent a memory leak:
	// var zero T
	// q.elements[0] = zero
	// ... if so, also confirm for Stack.Pop()
	q.elements = q.elements[1:]
	return result, true
}
