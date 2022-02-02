//go:build go1.18

package dictionary_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDictionary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generic Dictionary Suite")
}
