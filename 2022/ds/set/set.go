package ds

type Set[T comparable] map[T]bool

func NewSet[T comparable]() Set[T] {
  return make(map[T]bool)
}

func (s *Set[T]) Add(el T) bool {
  if !s.Contains(el) {
    (*s)[el] = true
    return true
  }
  return false
}

func (s *Set[T]) Remove(el T) bool {
  if s.Contains(el) {
    delete(*s, el)
    return true
  }
  return false
}

func (s *Set[T]) Contains(el T) bool {
  _, ok := (*s)[el]
  return ok
}

func (s *Set[T]) Size() int {
  return len(*s)
}

