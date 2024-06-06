package list

// 单向链表

type S[T any] struct {
	length int
	head   *sitem[T]
	tail   *sitem[T]
}

type sitem[T any] struct {
	data T
	next *sitem[T]
}

func NewS[T any]() *S[T] {
	return new(S[T])
}

func NewSfrom[T any](v T) *S[T] {
	s := new(S[T])
	n := new(sitem[T])
	n.data = v

	s.head = n

	s.tail = n
	s.length = 1
	return s

}

// 链表长度
func (s *S[T]) Len() int {
	return s.length
}

// 链表重置
func (s *S[T]) Reset() {
	s.length = 0
	s.head = nil
	s.tail = nil
}

// 链表头部
func (s *S[T]) First() T {

	return s.head.data
}

// 链表尾部
func (s *S[T]) Last() T {

	return s.tail.data
}

// 尾部追加
func (s *S[T]) Push(value T) {

	n := new(sitem[T])
	n.data = value
	if s.length > 0 {

		prev := s.tail
		prev.next = n

	} else {
		s.head = n
	}
	s.tail = n
	s.length++
}

// 头部删除
func (s *S[T]) Shift() T {

	if s.length <= 0 {
		panic("shift from empty list")
	}
	v := s.head.data

	if s.length != 1 {
		first := s.head
		s.head = first.next
		first.next = nil
	} else {
		s.head = nil
		s.tail = nil
	}
	s.length--

	return v
}

// 头部追加
func (s *S[T]) UnShift(value T) {
	n := new(sitem[T])
	n.data = value
	if s.length > 0 {
		n.next = s.head
		s.head = n

	} else {
		s.head = n
		s.tail = n
	}
	s.length++
}

// 链表数组
func (s *S[T]) Slice() []T {

	arr := make([]T, 0, s.length)

	for n := s.head; n != nil; n = n.next {
		arr = append(arr, n.data)
	}

	return arr
}

// f多为闭包函数
func (s *S[T]) ForEach(f func(T)) {

	for n := s.head; n != nil; n = n.next {
		f(n.data)
	}

}

func (s *S[T]) ForEachWithIndex(f func(T, int)) {
	k := 0
	for n := s.head; n != nil; n = n.next {
		f(n.data, k)
		k++
	}
}
