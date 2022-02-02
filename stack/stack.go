package stack

import "github.com/patrickhuber/go-collections/list"

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Length() int
}

type stack struct {
	list list.List
}

func New() Stack {
	return &stack{
		list: list.New(),
	}
}

func (s *stack) Push(item interface{}) {
	s.list.Append(item)
}

func (s *stack) Pop() interface{} {
	if s.list.Length() == 0 {
		return nil
	}
	item := s.list.Get(s.list.Length() - 1)
	s.list.RemoveAt(s.list.Length() - 1)
	return item
}

func (s *stack) Length() int {
	return s.list.Length()
}
