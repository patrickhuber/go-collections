//go:build go1.18

package stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/patrickhuber/go-collections/stack"
)

var _ = Describe("Stack", func() {
	It("can push", func() {
		s := stack.New()
		s.Push(1)
		Expect(s.Length()).To(Equal(1))
	})
	It("can pop", func() {
		s := stack.New()
		s.Push(1)
		Expect(s.Pop()).To(Equal(1))
	})
	It("can reverse", func() {
		s := stack.New()
		for i := 0; i < 5; i++ {
			s.Push(i)
		}
		for i := 4; i >= 0; i-- {
			Expect(s.Pop()).To(Equal(i))
		}
	})
})
