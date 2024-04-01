package collection

// Set is a generic non thread-safe set implementation.
type Set[T comparable] map[T]struct{}

// Add adds an element to the set.
func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

// Contains returns true if the set contains the element.
func (s Set[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}

// Remove removes an element from the set.
func (s Set[T]) Remove(element T) {
	delete(s, element)
}
