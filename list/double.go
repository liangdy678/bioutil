package list

// 双向链表

type D[T any] struct {
	length int
	head   *ditem[T]
	tail   *ditem[T]
}

type ditem[T any] struct {
	data T
	next *ditem[T]
	prev *ditem[T]
}

func NewD[T any]() *D[T] {
	return new(D[T])
}

func NewDfrom[T any](v T) *D[T] {
	d:= new(D[T])

	item := &ditem[T]{
		data: v,
	}
	d.head = item
	d.tail = item
	d.length=1
	return d
}



// 链表重置
func (s *D[T]) Reset() {
	s.length = 0
	s.head = nil
	s.tail = nil
}

// 链表长度
func (s *D[T]) Len() int {

	return s.length
}

// 链表头部
func (s *D[T]) First() T {

	return s.head.data
}

// 链表尾部
func (s *D[T]) Last() T {

	return s.tail.data
}

// 尾部追加
func (s *D[T]) Push(value T) {

	newtail := &ditem[T]{
		data: value,
	}

	if s.length > 0 {
		oldtail := s.tail
		oldtail.next = newtail
		newtail.prev = oldtail

	} else {
		s.head = newtail
	}

	s.tail = newtail

	s.length++
}

// 尾部删除
func (s *D[T]) Pop() T {

	if s.length <= 0 {
		panic("pop from empty list")
	}
	oldtail := s.tail
	v := oldtail.data
	newtail := oldtail.prev

	oldtail.prev = nil
	newtail.next = nil

	s.tail = newtail
	s.length--

	return v
}

// 头部追加

func (s *D[T]) UnShift(value T) {

	newfirst := &ditem[T]{
		data: value,
	}

	if s.length > 0 {

		oldfirst := s.head
		oldfirst.prev = newfirst
		newfirst.next = oldfirst

		s.head = newfirst

	} else {
		s.head = newfirst
		s.tail = newfirst
	}

	s.length++
}

// 头部删除
func (s *D[T]) Shift() T {

	if s.length <= 0 {
		panic("shift from empty list")

	}

	oldfirst := s.head

	v := oldfirst.data

	newfirst := oldfirst.next

	oldfirst.next = nil
	newfirst.prev = nil

	s.head = newfirst

	s.length--

	return v
}

// 双向链表数组
func (s *D[T]) Slice() []T {

	arr := make([]T, 0, s.length)

	for n := s.head; n != nil; n = n.next {
		arr = append(arr, n.data)
	}

	return arr
}

// 双向链表数组反转
func (s *D[T]) RevSlice() []T {

	arr := make([]T, 0, s.length)

	for n := s.tail; n != nil; n = n.prev {
		arr = append(arr, n.data)
	}
	return arr
}

// f多为闭包函数
func (s *D[T]) ForEach(f func(T)) {

	for n := s.head; n != nil; n = n.next {
		f(n.data)
	}

}
