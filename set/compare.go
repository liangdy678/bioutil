// set
package set

type Set[T comparable] map[T]void

func NewSet[T comparable](n int) Set[T] {
	return make(map[T]void, n)
}

func (s Set[T]) add(v T) {
	s[v] = void{}
}

func (s Set[T]) Add(v ...T) {

	for k := range v {
		s.add(v[k])
	}

}

func (s Set[T]) Discard(v T) {
	delete(s, v)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Clear() {
	for v := range s {
		delete(s, v)
	}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Intersect(o Set[T]) Set[T] { //交集
	r := NewSet[T](s.Len() + o.Len())

	if s.Len() < o.Len() {

		for v := range s {
			if o.Has(v) {
				r.Add(v)
			}
		}
		return r
	}

	for v := range o {
		if s.Has(v) {
			r.Add(v)
		}
	}

	return r

}

func (s Set[T]) Slice() []T {
	a := make([]T, 0, s.Len())
	for k := range s {
		a = append(a, k)
	}
	return a
}
