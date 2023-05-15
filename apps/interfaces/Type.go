package interfaces

type Dict[K comparable, V any] map[K]V
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](e ...T) Set[T] {
	s := Set[T]{}
	for _, v := range e {
		s.Add(v)
	}
	return s
}

func (s *Set[T]) Add(e ...T) {
	for _, v := range e {
		(*s)[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(e ...T) {
	for _, v := range e {
		delete(*s, v)
	}
}

func (s *Set[T]) Has(e T) bool {
	_, ok := (*s)[e]
	return ok
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Clear() {
	*s = Set[T]{}
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set[T]) ToSlice() []T {
	ret := make([]T, 0, s.Len())
	for k := range *s {
		ret = append(ret, k)
	}
	return ret
}
