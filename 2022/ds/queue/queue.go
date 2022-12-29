package ds

type Queue[T any] struct {
  queue []T
  size int
}

func NewQueue[T any]() *Queue[T] {
  return &Queue[T]{make([]T, 0), 0}
}

func (q *Queue[T]) Enqueue(el T) {
  q.queue = append(q.queue, el)
  q.size++
}

func (q *Queue[T]) Dequeue() T {
  defer func() {
    q.queue = q.queue[1:]
    q.size--
  }()
  return q.Top()
}

func (q *Queue[T]) Top() T {
  return q.queue[0]
}

func (q *Queue[T]) IsEmpty() bool {
  return q.size == 0
}

func (q *Queue[T]) GetQueue() []T {
  return q.queue
}

func (q *Queue[T]) GetSize() int {
  return q.size
}

