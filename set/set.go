package set

type Set[T comparable] struct {
	elements map[T]bool
}

func NewSet[T comparable](elems ...T) *Set[T] {
	set := &Set[T]{elements: make(map[T]bool)}

	for _, elem := range elems {
		set.Add(elem)
	}

	return set
}

func (set *Set[T]) Add(elem T) {
	set.elements[elem] = true
}

func (set *Set[T]) Remove(elem T) {
	delete(set.elements, elem)
}

func (set *Set[T]) Contains(elem T) bool {
	_, exists := set.elements[elem]
	return exists
}

func (set *Set[T]) Iterate() <-chan T {
	ch := make(chan T)
	go func() {
		for elem := range set.elements {
			ch <- elem
		}
		close(ch)
	}()
	return ch
}

func (set *Set[T]) Size() int {
	return len(set.elements)
}
