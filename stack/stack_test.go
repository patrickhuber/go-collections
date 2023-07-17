package stack_test

import (
	"testing"

	"github.com/patrickhuber/go-collections/stack"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Run("push", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		require.Equal(t, 1, s.Length())
	})
	t.Run("pop", func(t *testing.T) {

		s := stack.New()
		s.Push(1)
		require.Equal(t, 1, s.Pop())
	})
	t.Run("reverse", func(t *testing.T) {
		s := stack.New()
		for i := 0; i < 5; i++ {
			s.Push(i)
		}
		for i := 4; i >= 0; i-- {
			require.Equal(t, i, s.Pop())
		}
	})
}
