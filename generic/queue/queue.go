//go:build go1.18

package queue

import "github.com/patrickhuber/go-collections/generic/list"

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Length() int
}

type queue[T any] struct {
	list list.List[T]
}

func New[T any]() Queue[T] {
	return &queue[T]{
		list: list.New[T](),
	}
}

func (q *queue[T]) Enqueue(item T) {
	q.list.Insert(0, item)
}

func (q *queue[T]) Dequeue() T {
	var zero T
	if q.Length() == 0 {
		return zero
	}
	item := q.list.Get(q.list.Length() - 1)
	q.list.RemoveAt(q.list.Length() - 1)
	return item
}

func (q *queue[T]) Length() int {
	return q.list.Length()
}
