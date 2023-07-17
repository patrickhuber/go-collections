package list_test

import (
	"testing"

	"github.com/patrickhuber/go-collections/concurrent/list"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("can_create", func(t *testing.T) {
		list := list.New(1, 2, 3, 4)
		require.Equal(t, 4, list.Length())
	})
	t.Run("can append", func(t *testing.T) {
		list := list.New()
		list.Append(1)
		require.Equal(t, 1, list.Length())
		require.Equal(t, 1, list.Get(0))
	})
	t.Run("can remove", func(t *testing.T) {
		list := list.New(1, 3, 4, 5)
		list.Remove(4)
		require.Equal(t, 3, list.Length())
		require.Equal(t, 1, list.Get(0))
		require.Equal(t, 3, list.Get(1))
		require.Equal(t, 5, list.Get(2))
	})
	t.Run("can clear", func(t *testing.T) {
		list := list.New(1, 3, 4, 5)
		list.Clear()
		require.Equal(t, 0, list.Length())
	})

	t.Run("contains_item", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5, 6)
		require.True(t, list.Contains(3))
	})

	t.Run("missing_item", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5, 6)
		require.False(t, list.Contains(10))
	})

	t.Run("index_of", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5)
		index := list.IndexOf(2)
		require.Equal(t, 1, index)
	})

	t.Run("for_each", func(t *testing.T) {
		list := list.New(1, 2, 3, 4, 5)
		sum := 0
		list.ForEach(func(item any) {
			sum += item.(int)
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
}
