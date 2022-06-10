//go:build go1.18

package list

import "github.com/patrickhuber/go-collections/object"

type ReadOnlyList[T any] interface {
	Get(index int) T
	Length() int
	IndexOf(item T) int
	Contains(item T) bool
	ForEach(func(T))
}

type List[T any] interface {
	ReadOnlyList[T]
	Set(index int, item T)
	Append(item T)
	Insert(index int, item T)
	Remove(T)
	RemoveAt(index int)
	Clear()
}

type list[T any] struct {
	items []T
}

func New[T any](items ...T) List[T] {
	return &list[T]{
		items: items,
	}
}

func (l *list[T]) Length() int {
	return len(l.items)
}

func (l *list[T]) Get(index int) T {
	return l.items[index]
}

func (l *list[T]) Set(index int, item T) {
	l.items[index] = item
}

func (l *list[T]) Append(item T) {
	l.items = append(l.items, item)
}

func (l *list[T]) Insert(index int, item T) {
	if l.Length() == index {
		l.Append(item)
		return
	}
	l.items = append(l.items[:index+1], l.items[index:]...)
	l.items[index] = item
}

func (l *list[T]) IndexOf(item T) int {
	for index, value := range l.items {
		if object.Equal(value, item) {
			return index
		}
	}
	return -1
}

func (l *list[T]) Contains(item T) bool {
	return l.IndexOf(item) >= 0
}

func (l *list[T]) Remove(item T) {
	index := l.IndexOf(item)
	l.RemoveAt(index)
}

func (l *list[T]) RemoveAt(index int) {
	var zero T
	copy(l.items[index:], l.items[index+1:]) // Shift a[i+1:] left one index.
	l.items[len(l.items)-1] = zero           // Erase last element (write zero value).
	l.items = l.items[:len(l.items)-1]       // Truncate slice.
}

func (l *list[T]) Clear() {
	l.items = nil
}

func (l *list[T]) ForEach(action func(item T)) {
	for _, item := range l.items {
		action(item)
	}
}
