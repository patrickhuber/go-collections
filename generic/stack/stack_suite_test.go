//go:build go1.18

package stack_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generic Stack Suite")
}
