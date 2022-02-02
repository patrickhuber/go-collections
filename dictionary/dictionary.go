package dictionary

type Dictionary interface {
	Get(key interface{}) interface{}
	Lookup(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{})
	Length() int
	Clear()
	Keys() []interface{}
	Values() []interface{}
}
type dictionary struct {
	values map[interface{}]interface{}
}

func New() Dictionary {
	return &dictionary{
		values: map[interface{}]interface{}{},
	}
}

func (d *dictionary) Get(key interface{}) interface{} {
	return d.values[key]
}

func (d *dictionary) Lookup(key interface{}) (interface{}, bool) {
	value, ok := d.values[key]
	return value, ok
}

func (d *dictionary) Set(key interface{}, value interface{}) {
	d.values[key] = value
}

func (d *dictionary) Length() int {
	return len(d.values)
}

func (d *dictionary) Clear() {
	for k := range d.values {
		delete(d.values, k)
	}
}

func (d *dictionary) Keys() []interface{} {
	keys := []interface{}{}
	for k := range d.values {
		keys = append(keys, k)
	}
	return keys
}

func (d *dictionary) Values() []interface{} {
	values := []interface{}{}
	for _, v := range d.values {
		values = append(values, v)
	}
	return values
}
