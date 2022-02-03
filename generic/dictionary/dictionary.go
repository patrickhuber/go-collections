//go:build go1.18

package dictionary

type ReadOnlyDictionary[TKey comparable, TValue any] interface {
	Get(key TKey) TValue
	Lookup(key TKey) (TValue, bool)
	Length() int
	Keys() []TKey
	Values() []TValue
}

type Dictionary[TKey comparable, TValue any] interface {
	ReadOnlyDictionary[TKey, TValue]
	Set(key TKey, value TValue)
	Clear()
	Remove(key TKey)
}

type dictionary[TKey comparable, TValue any] struct {
	items map[TKey]TValue
}

func New[TKey comparable, TValue any]() Dictionary[TKey, TValue] {
	return &dictionary[TKey, TValue]{
		items: map[TKey]TValue{},
	}
}

func NewWithMap[TKey comparable, TValue any](m map[TKey]TValue) Dictionary[TKey, TValue] {
	cpy := map[TKey]TValue{}
	for k, v := range m {
		cpy[k] = v
	}
	return &dictionary[TKey, TValue]{
		items: cpy,
	}
}

func (d *dictionary[TKey, TValue]) Get(key TKey) TValue {
	return d.items[key]
}

func (d *dictionary[TKey, TValue]) Lookup(key TKey) (TValue, bool) {
	value, ok := d.items[key]
	return value, ok
}

func (d *dictionary[TKey, TValue]) Set(key TKey, value TValue) {
	d.items[key] = value
}

func (d *dictionary[TKey, TValue]) Remove(key TKey) {
	delete(d.items, key)
}

func (d *dictionary[TKey, TValue]) Length() int {
	return len(d.items)
}

func (d *dictionary[TKey, TValue]) Clear() {
	for k := range d.items {
		delete(d.items, k)
	}
}

func (d *dictionary[TKey, TValue]) Keys() []TKey {
	keys := []TKey{}
	for k := range d.items {
		keys = append(keys, k)
	}
	return keys
}

func (d *dictionary[TKey, TValue]) Values() []TValue {
	values := []TValue{}
	for _, v := range d.items {
		values = append(values, v)
	}
	return values
}
