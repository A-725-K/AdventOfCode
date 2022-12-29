package ds

type Stack[T any] struct {
  data []T
  size int
}

func NewStack[T any]() *Stack[T] {
  return &Stack[T]{make([]T, 0), 0}
}

func (s *Stack[T]) Push(e T) {
  s.data = append(s.data, e)
  s.size++
}

// func (s *Stack[T]) PopLeft() T {
//   defer func() {
//     s.size--
//     s.data = s.data[1:]
//   }()
//   return s.data[0]
// }

func (s *Stack[T]) Pop() T {
  defer func() {
    s.size--
    s.data = s.data[:s.size]
  }()
  return s.Top()
}

func (s *Stack[T]) Top() T {
  return s.data[s.size-1]
}

func (s *Stack[T]) IsEmpty() bool {
  return s.size == 0
}

func (s *Stack[T]) GetStack() []T {
  return s.data
}

func (s *Stack[T]) GetSize() int {
  return s.size
}

