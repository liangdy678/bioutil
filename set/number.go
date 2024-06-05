// set
package set

type void = struct{}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Num[T Number] struct {
	m     map[T]void
	min   T
	max   T
	empty bool
}

func UseNumSet[T Number](n int) *Num[T] {
	return &Num[T]{m: make(map[T]void, n)}
}

func (s *Num[T]) Len() int {
	return len(s.m)
}

func (s *Num[T]) Add(v T) {

	if s.Has(v) {
		return
	}

	s.m[v] = void{}

	if s.empty {
		s.min = v
		s.max = v
		s.empty = false
		return
	}

	if v > s.max {
		s.max = v
		return
	}

	if v < s.min {
		s.min = v
	}

}

func (s *Num[T]) has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Num[T]) Has(v T) bool {
	return s.has(v)
}

func (s *Num[T]) Max() (T, bool) {
	ok := !s.empty
	return s.max, ok
}

func (s *Num[T]) Min() (T, bool) {
	ok := !s.empty

	return s.min, ok
}

func (s *Num[T]) Discard(v T) {

	if !s.has(v) {
		return
	}

	delete(s.m, v)

	if len(s.m) <= 0 {
		s.empty = true
		s.min = 0
		s.max = 0
		return
	}

	if s.min < v && v < s.max {
		return
	}

	//recalculate min and max
	var max, min T

	fst := true
	for v := range s.m {

		if fst {
			//赋值
			max, min = v, v
			fst = false
			continue
		}

		if v > max {
			max = v
			continue
		}

		if v < min {
			min = v
		}
	}

	s.max, s.min = max, min
}

func (s *Num[T]) Slice() []T {
	a := make([]T, 0, s.Len())
	for k := range s.m {
		a = append(a, k)
	}
	return a
}
