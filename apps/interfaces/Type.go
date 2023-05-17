package interfaces

type Dict[K comparable, V any] map[K]V
type Set[T comparable] map[T]struct{}

func NewDict[K comparable, V any]() Dict[K, V] {
	return Dict[K, V]{}
}

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

func (d *Dict[K, V]) Get(k K) (V, bool) {
	v, ok := (*d)[k]
	return v, ok
}

func (d *Dict[K, V]) Add(k K, v V) *Dict[K, V] {
	(*d)[k] = v
	return d
}

func (d *Dict[K, V]) Remove(e ...K) *Dict[K, V] {
	for _, v := range e {
		delete(*d, v)
	}
	return d
}

func (d *Dict[K, V]) Has(e K) bool {
	_, ok := (*d)[e]
	return ok
}

func (d *Dict[K, V]) Len() int {
	return len(*d)
}

func (d *Dict[K, V]) Clear() {
	*d = Dict[K, V]{}
}

func (d *Dict[K, V]) IsEmpty() bool {
	return d.Len() == 0
}
