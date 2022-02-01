package queue

import "github.com/patrickhuber/go-collections/list"

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Length() int
}

type queue struct {
	list list.List
}

func New() Queue {
	return &queue{
		list: list.New(),
	}
}

func (q *queue) Enqueue(item interface{}) {
	q.list.Insert(0, item)
}

func (q *queue) Dequeue() interface{} {
	item := q.list.Get(q.list.Length() - 1)
	q.list.RemoveAt(q.list.Length() - 1)
	return item
}

func (q *queue) Length() int {
	return q.list.Length()
}
