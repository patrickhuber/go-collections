//go:build go1.18

package stack

import "github.com/patrickhuber/go-collections/generic/list"

type Stack[T any] interface {
	Push(T)
	Pop() T
	Length() int
}

type stack[T any] struct {
	list list.List[T]
}

func New[T any]() Stack[T] {
	return &stack[T]{
		list: list.New[T](),
	}
}

func (s *stack[T]) Push(item T) {
	s.list.Append(item)
}

func (s *stack[T]) Pop() T {
	var zero T
	if s.list.Length() == 0 {
		return zero
	}
	item := s.list.Get(s.list.Length() - 1)
	s.list.RemoveAt(s.list.Length() - 1)
	return item
}

func (s *stack[T]) Length() int {
	return s.list.Length()
}
