package list

import (
	"sync"

	"github.com/patrickhuber/go-collections/list"
)

type concurrentList struct {
	innerList list.List
	mut       sync.RWMutex
}

func New(values ...interface{}) list.List {
	list := list.New(values...)
	return &concurrentList{
		innerList: list,
	}
}

func (l *concurrentList) Get(index int) interface{} {
	l.mut.RLock()
	defer l.mut.RUnlock()
	return l.innerList.Get(index)
}

func (l *concurrentList) Set(index int, value interface{}) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.innerList.Set(index, value)
}

func (l *concurrentList) Insert(index int, value interface{}) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.innerList.Insert(index, value)
}

func (l *concurrentList) Append(value interface{}) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.innerList.Append(value)
}

func (l *concurrentList) RemoveAt(index int) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.innerList.RemoveAt(index)
}

func (l *concurrentList) Remove(value interface{}) {
	index := l.IndexOf(value)
	l.RemoveAt(index)
}

func (l *concurrentList) Length() int {
	l.mut.RLock()
	defer l.mut.RUnlock()
	return l.innerList.Length()
}

func (l *concurrentList) Contains(value interface{}) bool {
	// no mutex needed here because IndexOf takes care of the lock
	return l.IndexOf(value) >= 0
}

func (l *concurrentList) IndexOf(value interface{}) int {
	l.mut.RLock()
	defer l.mut.RUnlock()

	return l.innerList.IndexOf(value)
}

func (l *concurrentList) ForEach(delegate func(item interface{})) {
	l.mut.RLock()
	defer l.mut.RUnlock()

	for i := 0; i < l.Length(); i++ {
		value := l.innerList.Get(i)
		delegate(value)
	}
}

func (l *concurrentList) Clear() {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.innerList.Clear()
}
