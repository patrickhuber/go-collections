//go:build go1.18

package list_test

import (
	"github.com/patrickhuber/go-collections/generic/list"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {
	It("can create", func() {
		list := list.New(1, 2, 3, 4)
		Expect(list.Length()).To(Equal(4))
	})
	It("can append", func() {
		list := list.New[int]()
		list.Append(1)
		Expect(list.Length()).To(Equal(1))
		Expect(list.Get(0)).To(Equal(1))
	})
	It("can remove", func() {
		list := list.New(1, 3, 4, 5)
		list.Remove(4)
		Expect(list.Length()).To(Equal(3))
		Expect(list.Get(0)).To(Equal(1))
		Expect(list.Get(1)).To(Equal(3))
		Expect(list.Get(2)).To(Equal(5))
	})
	It("can clear", func() {
		list := list.New(1, 3, 4, 5)
		list.Clear()
		Expect(list.Length()).To(Equal(0))
	})
	When("item in list", func() {
		It("returns true", func() {
			list := list.New(1, 2, 3, 4, 5, 6)
			Expect(list.Contains(3)).To(BeTrue())
		})
	})
	When("item missing", func() {
		It("returns false", func() {
			list := list.New(1, 2, 3, 4, 5, 6)
			Expect(list.Contains(10)).To(BeFalse())
		})
	})
	It("can return indexof", func() {
		list := list.New(1, 2, 3, 4, 5)
		index := list.IndexOf(2)
		Expect(index).To(Equal(1))
	})
	It("can run foreach", func() {
		list := list.New(1, 2, 3, 4, 5)
		sum := 0
		list.ForEach(func(item int) {
			sum += item
		})
		Expect(sum).To(Equal(15))
	})
	It("can get", func() {
		list := list.New(1, 3, 4, 5, 6)
		item := list.Get(3)
		Expect(item).To(Equal(5))
	})
	It("can set", func() {
		list := list.New(1, 3, 4, 5, 6)
		list.Set(0, 10)
		item := list.Get(0)
		Expect(item).To(Equal(10))
	})
})
