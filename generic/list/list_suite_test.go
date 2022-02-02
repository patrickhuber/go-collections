//go:build go1.18

package list_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generic List Suite")
}
