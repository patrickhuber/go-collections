package queue_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/patrickhuber/go-collections/queue"
)

func TestQueue(t *testing.T) {
	t.Run("enqueue", func(t *testing.T) {
		q := queue.New()
		q.Enqueue("test")
		require.Equal(t, 1, q.Length())
	})

	t.Run("dequeue", func(t *testing.T) {
		q := queue.New()
		q.Enqueue("test")
		item := q.Dequeue()
		require.Equal(t, "test", item)

	})
	t.Run("fifo", func(t *testing.T) {
		q := queue.New()
		for i := 0; i < 5; i++ {
			q.Enqueue(i)
		}
		for i := 0; i < 5; i++ {
			value := q.Dequeue()
			require.Equal(t, i, value)
		}
	})
}
