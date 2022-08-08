package ratelimiter

type Queue[T any] interface {
	Enqueue(element T)
	Dequeue() (T, bool)
	GetSize() int
}

type queue[T any] struct {
	bucket []T
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{
		bucket: []T{},
	}
}

func (q *queue[T]) tryDequeue() (T, bool) {
	if len(q.bucket) == 0 {
		var dummy T
		return dummy, false
	}
	value := q.bucket[0]
	var zero T
	q.bucket[0] = zero // Avoid memory leak
	q.bucket = q.bucket[1:]
	return value, true
}

func (q *queue[T]) Enqueue(element T) {
	q.bucket = append(q.bucket, element)
}

func (q *queue[T]) Dequeue() (T, bool) {
	if len(q.bucket) == 0 {
		var dummy T
		return dummy, false
	}

	value := q.bucket[0]
	var zero T
	q.bucket[0] = zero // Avoid memory leak
	q.bucket = q.bucket[1:]

	return value, true
}

func (q *queue[T]) GetSize() int {
	return len(q.bucket)
}
