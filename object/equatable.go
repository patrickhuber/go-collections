package object

import "reflect"

type Equatable interface {
	Equal(any) bool
}

func Equal(left any, right any) bool {
	eq, ok := left.(Equatable)
	if ok {
		return eq.Equal(right)
	}
	return reflect.DeepEqual(left, right)
}
