package collection

// Stack is a generic non thread-safe LIFO stack implementation.
type Stack[T any] struct {
	elements []T
}

// NewStack creates a new stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// NewStackWithCapacity creates a new stack with the specified initial capacity.
func NewStackWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0, capacity),
	}
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.elements = s.elements[:0]
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.elements)
}

// Cap returns the capacity of the stack.
func (s *Stack[T]) Cap() int {
	return cap(s.elements)
}

// Peek returns the top element of the stack (without removing it)
// and a boolean indicating if the operation was successful.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var result T
		return result, false
	}

	return s.elements[len(s.elements)-1], true
}

// Pop removes and returns the top element of the stack and a boolean indicating
// if the operation was successful.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var result T
		return result, false
	}

	endIndex := len(s.elements) - 1
	result := s.elements[endIndex]
	s.elements = s.elements[:endIndex]
	return result, true
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}
