package dictionary_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/patrickhuber/go-collections/dictionary"
)

var _ = Describe("Dictionary", func() {
	It("can set", func() {
		d := dictionary.New()
		d.Set("test", "1")
		Expect(d.Length()).To(Equal(1))
	})
	It("can get", func() {
		d := dictionary.New()
		d.Set("test", "1")
		value := d.Get("test")
		Expect(value).To(Equal("1"))
	})
	It("can list keys", func() {
		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		keys := d.Keys()
		Expect(len(keys)).To(Equal(2))
		Expect(keys[0]).To(Equal("test"))
		Expect(keys[1]).To(Equal("other"))
	})
	It("can list values", func() {
		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		values := d.Values()
		Expect(len(values)).To(Equal(2))
		Expect(values[0]).To(Equal("1"))
		Expect(values[1]).To(Equal("2"))
	})
	It("can clear", func() {
		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		d.Clear()
		Expect(d.Length()).To(Equal(0))
	})
	Describe("Lookup", func() {
		When("missing", func() {
			It("returns false", func() {
				d := dictionary.New()
				d.Set("test", "1")
				d.Set("other", "2")
				_, ok := d.Lookup("hello")
				Expect(ok).To(BeFalse())
			})
		})
		When("present", func() {
			It("can lookup", func() {
				d := dictionary.New()
				d.Set("test", "1")
				d.Set("other", "2")
				value, ok := d.Lookup("test")
				Expect(ok).To(BeTrue())
				Expect(value).To(Equal("1"))
			})
		})
	})
	It("can remove", func() {
		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		d.Remove("other")
		Expect(d.Length()).To(Equal(1))
	})
})
