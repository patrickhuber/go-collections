package queue_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/patrickhuber/go-collections/queue"
)

var _ = Describe("Queue", func() {
	It("can enqueue", func() {
		q := queue.New()
		q.Enqueue("test")
		Expect(q.Length()).To(Equal(1))
	})
	It("can dequeue", func() {
		q := queue.New()
		q.Enqueue("test")
		item := q.Dequeue()
		Expect(item).To(Equal("test"))
	})
	It("can fifo", func() {
		q := queue.New()
		for i := 0; i < 5; i++ {
			q.Enqueue(i)
		}
		for i := 0; i < 5; i++ {
			value := q.Dequeue()
			Expect(value).To(Equal(i))
		}
	})
})
