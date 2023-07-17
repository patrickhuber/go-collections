//go:build go1.18

package list_test

import (
	"testing"

	"github.com/patrickhuber/go-collections/generic/list"
	"github.com/stretchr/testify/require"
)

type S struct {
	A string
}

func (s *S) FuncA() string {
	return s.A
}

func (s *S) Equal(other any) bool {
	if s == nil && other == nil {
		return true
	}
	if other == nil {
		return false
	}
	right, ok := other.(SI)
	if !ok {
		return false
	}
	return right.FuncA() == s.A
}

type SI interface {
	FuncA() string
}

func NewSI(a string) SI {
	return &S{
		A: a,
	}
}

func TestList(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		list := list.New(1, 2, 3, 4)
		require.Equal(t, 4, list.Length())
	})
	t.Run("append", func(t *testing.T) {
		list := list.New[int]()
		list.Append(1)
		require.Equal(t, 1, list.Length())
		require.Equal(t, 1, list.Get(0))
	})
	t.Run("remove", func(t *testing.T) {
		list := list.New(1, 3, 4, 5)
		list.Remove(4)
		require.Equal(t, 3, list.Length())
		require.Equal(t, 1, list.Get(0))
		require.Equal(t, 3, list.Get(1))
		require.Equal(t, 5, list.Get(2))
	})
	t.Run("clear", func(t *testing.T) {
		list := list.New(1, 3, 4, 5)
		list.Clear()
		require.Equal(t, 0, list.Length())
	})
	t.Run("contains", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5, 6)
		require.True(t, list.Contains(3))
	})
	t.Run("missing", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5, 6)
		require.False(t, list.Contains(10))
	})
	t.Run("indexof", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5)
		index := list.IndexOf(2)
		require.Equal(t, 1, index)
	})
	t.Run("foreach", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5)
		sum := 0
		list.ForEach(func(item int) {
			sum += item
		})
		require.Equal(t, 15, sum)
	})
	t.Run("get", func(t *testing.T) {

		list := list.New(1, 3, 4, 5, 6)
		item := list.Get(3)
		require.Equal(t, 5, item)
	})
	t.Run("set", func(t *testing.T) {

		list := list.New(1, 3, 4, 5, 6)
		list.Set(0, 10)
		item := list.Get(0)
		require.Equal(t, 10, item)
	})
	t.Run("new", func(t *testing.T) {
		list := list.New(NewSI("test"))
		require.NotNil(t, list)
		require.Equal(t, 1, list.Length())
		require.NotNil(t, list.Get(0))
	})
	t.Run("equality", func(t *testing.T) {
		list := list.New(NewSI("test"))
		other := NewSI("test")
		require.True(t, list.Contains(other))
	})
}
