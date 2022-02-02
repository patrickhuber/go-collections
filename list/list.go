package list

type ReadOnlyList interface {
	Get(index int) interface{}
	Contains(item interface{}) bool
	IndexOf(item interface{}) int
	Length() int
	ForEach(func(interface{}))
}

type List interface {
	ReadOnlyList
	Set(index int, value interface{})
	Append(value interface{})
	Insert(index int, value interface{})
	Remove(item interface{})
	RemoveAt(index int)
	Clear()
}

type list struct {
	items []interface{}
}

func New(items ...interface{}) List {
	return &list{
		items: items,
	}
}

func NewReadOnly(items ...interface{}) ReadOnlyList {
	return New(items...)
}

func (l *list) Length() int {
	return len(l.items)
}

func (l *list) Get(index int) interface{} {
	return l.items[index]
}

func (l *list) Set(index int, value interface{}) {
	l.items[index] = value
}

func (l *list) Append(item interface{}) {
	l.items = append(l.items, item)
}

func (l *list) Insert(index int, item interface{}) {
	if l.Length() == index {
		l.Append(item)
		return
	}
	l.items = append(l.items[:index+1], l.items[index:]...)
	l.items[index] = item
}

func (l *list) IndexOf(item interface{}) int {
	for index, value := range l.items {
		if value == item {
			return index
		}
	}
	return -1
}

func (l *list) Contains(item interface{}) bool {
	return l.IndexOf(item) >= 0
}

func (l *list) Remove(item interface{}) {
	index := l.IndexOf(item)
	l.RemoveAt(index)
}

func (l *list) RemoveAt(index int) {
	copy(l.items[index:], l.items[index+1:]) // Shift a[i+1:] left one index.
	l.items[len(l.items)-1] = nil            // Erase last element (write zero value).
	l.items = l.items[:len(l.items)-1]       // Truncate slice.
}

func (l *list) Clear() {
	l.items = nil
}

func (l *list) ForEach(action func(item interface{})) {
	for _, item := range l.items {
		action(item)
	}
}
