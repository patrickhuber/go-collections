//go:build go1.18

package stack

import "github.com/patrickhuber/go-collections/generic/list"

type Stack[T any] interface {
	// Push pushes a item onto the stack
	Push(T)
	// Pop removes an item from the top of the stack and returns the item.
	// If there is no item on the top of the stack, pop returns the default value for T.
	Pop() T
	// Length returns the number of items in the stack
	Length() int
	// Peek returns the top of the stack and true if not empty.
	// If the stack is empty, peek returns the default value for T and false
	Peek() (T, bool)
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

func (s *stack[T]) Peek() (T, bool) {
	if s.list.Length() > 0 {
		return s.list.Get(0), true
	}
	var result T
	return result, false
}
